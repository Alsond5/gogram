package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Alsond5/gogram"
	"github.com/Alsond5/gogram/models"
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

func helloHandler(ctx context.Context, bot *gogram.Bot, update *models.Update) {
	bot.SendMessageHelper(ctx, fmt.Sprintf("Hello, %s", update.Message.From.FirstName), update.Message.Chat.Id)
}
