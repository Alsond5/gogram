package gogram

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Alsond5/gogram/models"
)

func defaultHandler(_ context.Context, _ *Bot, upd *models.Update) {
	b, err := json.Marshal(upd)
	if err != nil {
		return
	}

	fmt.Printf("[UPDATE] %s\n", b)
}

func wrapMiddlewares(h HandlerFunc, m ...Middleware) HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func (b *Bot) processUpdate(ctx context.Context, upd *models.Update) {
	handler := b.findHandler(upd)

	w := wrapMiddlewares(handler, b.middlewares...)
	go w(ctx, b, upd)
}

func (b *Bot) findHandler(upd *models.Update) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	for _, h := range b.handlers {
		if h.match(upd) {
			return h.handler
		}
	}

	return defaultHandler
}

func (b *Bot) Command(upd *models.Update) string {
	if upd.Message == nil || upd.Message.Entities == nil {
		return ""
	}

	for _, entity := range upd.Message.Entities {
		if entity.Type == "bot_command" {
			cmd := upd.Message.Text[entity.Offset : entity.Offset+entity.Length]
			return cmd
		}
	}

	return ""
}
