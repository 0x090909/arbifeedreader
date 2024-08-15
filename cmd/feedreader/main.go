package main

import (
	"github.com/0x090909/arbifeedreader/arbitrumtypes"
	"github.com/0x090909/arbifeedreader/feedreader"
	log "github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "01-02|15:04:05.999"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	f := feedreader.NewFeedReader()
	f.Run(func(tx *arbitrumtypes.Transaction) {
		log.WithFields(log.Fields{"Hash": tx.Hash()}).Info("New Transaction!")
		log.WithFields(log.Fields{"To": tx.To()}).Info("New Transaction!")
		//log.WithFields(log.Fields{"Data": tx.Data()}).Info("New Transaction!")
	})
}
