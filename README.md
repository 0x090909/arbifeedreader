# ArbitrumFeedReader

ArbitrumFeedReader is a Go library that allows you to subscribe to and parse transactions coming from the Arbitrum One sequencer.

## Installation

To install ArbitrumFeedReader, use `go get`:

```
go get github.com/0x090909/arbifeedreader
```

## Usage

Here's a basic example of how to use ArbitrumFeedReader:

```go
package main

import (
    "github.com/0x090909/arbifeedreader"
    "github.com/sirupsen/logrus"
)

func main() {
    f := arbifeedreader.NewFeedReader()
    f.Run(func(tx *arbitrumtypes.Transaction) {
        logrus.WithFields(logrus.Fields{"Hash": tx.Hash()}).Info("New Transaction!")
        logrus.WithFields(logrus.Fields{"To": tx.To()}).Info("New Transaction!")
        // Uncomment the following line to log transaction data
        // logrus.WithFields(logrus.Fields{"Data": tx.Data()}).Info("New Transaction!")
    })
}
```

This example creates a new FeedReader, starts it, and logs the hash and recipient address of each new transaction.

## API Reference

### `NewFeedReader()`

Creates and returns a new FeedReader instance.

### `FeedReader.Run(callback func(*arbitrumtypes.Transaction))`

Starts the FeedReader and calls the provided callback function for each new transaction.

The callback function receives a pointer to an `arbitrumtypes.Transaction` object, which represents a transaction on the Arbitrum One network.

## Dependencies

This library depends on:

- `github.com/sirupsen/logrus` for logging

Make sure to import and initialize these dependencies in your project.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT
