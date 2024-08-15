package arbostypes

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

var uniquifyingPrefix = []byte("Arbitrum Nitro Feed:")

type MessageWithMetadata struct {
	Message             *L1IncomingMessage `json:"message"`
	DelayedMessagesRead uint64             `json:"delayedMessagesRead"`
}

type MessageWithMetadataAndBlockHash struct {
	MessageWithMeta MessageWithMetadata
	BlockHash       *common.Hash
}

var EmptyTestMessageWithMetadata = MessageWithMetadata{
	Message: &EmptyTestIncomingMessage,
}

// TestMessageWithMetadataAndRequestId message signature is only verified if requestId defined
var TestMessageWithMetadataAndRequestId = MessageWithMetadata{
	Message: &TestIncomingMessageWithRequestId,
}

type InboxMultiplexer interface {
	Pop(context.Context) (*MessageWithMetadata, error)
	DelayedMessagesRead() uint64
}
