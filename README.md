Telegram API in Go
===============
This library supports APIS from the `api.telegram.org` REST calls. At the moment, this library only supports text based APIs.
Support for webhook and multimedia (including telegram stickers) will be added in subsequent updates.

## Installing
### *go get*
    $ go get -u github.com/chenlu-chua/go-telegram

## Example
### Getting updates
```golang
package main

import (
	"context"
	"fmt"
	"github.com/ccl17/go-telegram"
)

func main() {
	
	// Create a telegram bot client with default options
	bot := telegram.New("YOUR_SECRET_BOT_TOKEN")
	
	updatesCh := make(chan telegram.Update, 100)
	killCh := make(chan struct{})
	
	go func() {
		
		opts := &telegram.GetUpdatesOptions{
			// Updates from telegram API server will return updates with update_id
			// greater than the offset in a running order. This number is used as
			// a means to mark an update as handled.
			Offset: 693786789,
			Limit: 100,
			// Update request timeout for long polling.
			Timeout: 5,
		}
		
		for {
			updates, err := bot.GetUpdates(context.Background(), opts)
			if err != nil {
				fmt.Printf("%s\n", err)
				continue
			}
			
			for _, update := range updates {
				updatesCh <- update
			}
			
			select {
			case <- killCh:
				return
			default:
			}
		}
	}()
}
```