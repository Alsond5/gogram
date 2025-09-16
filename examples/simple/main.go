package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Alsond5/gogram"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token := os.Getenv("BOT_TOKEN")

	b, err := gogram.NewBot(token)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(gogram.HandlerTypeMessageText, "hello", gogram.MatchTypeCommandStartOnly, helloHandler)

	b.Start(ctx)
}

func helloHandler(c *gogram.Context) error {
	_, err := c.Bot.SendMessageHelper(c.Ctx, fmt.Sprintf("Hello, %s", c.Update.Message.From.FirstName), c.Update.Message.Chat.Id)

	return err
}
