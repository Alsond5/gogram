package gogram

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Alsond5/gogram/models"
)

func (b *Bot) startPooling(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		params := &models.GetUpdatesRequest{
			Offset: atomic.LoadInt64(&b.lastUpdateID) + 1,
		}

		var updates []models.Update
		updates, err := b.GetUpdates(ctx, params)
		if err != nil {
			fmt.Println("getUpdates error:", err)
			time.Sleep(time.Second)

			continue
		}

		for _, u := range updates {
			if u.UpdateId > b.lastUpdateID {
				atomic.StoreInt64(&b.lastUpdateID, u.UpdateId)
				b.updates <- &u
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}
