package gogram

import (
	"context"

	"github.com/Alsond5/gogram/models"
)

type Context struct {
	Ctx      context.Context
	Bot      *Bot
	Update   *models.Update
	handlers []Handler
	index    int
}

func (c *Context) Next() error {
	c.index++

	if c.index >= len(c.handlers) {
		return nil
	}

	return c.handlers[c.index](c)
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}
