package gogram

import (
	"context"
	"sync"
)

func (b *Bot) handleUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case upd := <-b.updates:
			b.processUpdate(ctx, upd)
		}
	}
}
