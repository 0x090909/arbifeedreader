package feedreader

import (
	"encoding/json"
	"github.com/0x090909/arbifeedreader/arbitrumtypes"
	"github.com/0x090909/arbifeedreader/arbos"
	"github.com/0x090909/arbifeedreader/broadcaster/message"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type FeedReader struct {
	endpoint string
}

func NewFeedReader() *FeedReader {
	return &FeedReader{endpoint: "wss://arb1.arbitrum.io/feed"}
}

func (f *FeedReader) Run(callback func(transaction *arbitrumtypes.Transaction)) {

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(f.endpoint, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing WebSocket connection:", err)
		}
	}(conn)

	for {
		// Read message from the WebSocket
		_, messageWSS, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		// Decode the JSON message into a BroadcastMessage struct
		var broadcastMsg message.BroadcastMessage
		err = json.Unmarshal(messageWSS, &broadcastMsg)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			continue
		}

		// Process the decoded message
		for _, m := range broadcastMsg.Messages {
			if m.Message.Message != nil {
				txes, errorParse := arbos.ParseL2Transactions(m.Message.Message, big.NewInt(42161))
				if errorParse != nil {
					log.Println("Error parsing txes:", errorParse)

				}
				if len(txes) > 0 {
					for _, tx := range txes {
						if tx.To() != nil {
							callback(tx)
						}

					}
				}
			}

		}
		//fmt.Printf("Received message - Type: %s, Content: %s\n", broadcastMsg.Messages, broadcastMsg.ConfirmedSequenceNumberMessage)
	}
}
