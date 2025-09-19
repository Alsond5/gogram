package gogram

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Alsond5/gogram/models"
)

func defaultHandler(c *Context) error {
	b, err := json.Marshal(c.Update)
	if err != nil {
		return err
	}

	fmt.Printf("[UPDATE] %s\n", b)

	return nil
}

func (b *Bot) processUpdate(ctx context.Context, upd *models.Update) {
	handlers := b.findHandler(upd)

	c := &Context{
		Ctx:          ctx,
		Bot:          b,
		Update:       upd,
		handlers:     handlers,
		handlerCount: len(handlers),
		index:        -1,
	}

	go c.Next()
}

func (b *Bot) findHandler(upd *models.Update) []Handler {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	for _, h := range b.handlers {
		if h.match(upd) {
			middlewareCount := h.middlewareSnapshot

			if middlewareCount == 0 {
				return h.routes
			}

			middlewares := b.getMiddlewares(middlewareCount)

			totalLen := middlewareCount + len(h.routes)
			routes := make([]Handler, totalLen)

			copy(routes, middlewares)
			copy(routes[middlewareCount:], h.routes)

			return routes
		}
	}

	return []Handler{defaultHandler}
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
