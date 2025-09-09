package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Alsond5/gogram"
	"github.com/Alsond5/gogram/models"
)

var botCommands = []models.BotCommand{
	{
		Command:     "/start",
		Description: "Start the bot",
	},
	{
		Command:     "/help",
		Description: "Show available commands and usage information",
	},
	{
		Command:     "/hello",
		Description: "Send a friendly greeting",
	},
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token := os.Getenv("BOT_TOKEN")

	b, err := gogram.NewBot(token, gogram.SetWorkers(3), gogram.SetMiddlewares(printMessageMiddleware))
	if err != nil {
		fmt.Println(err)
	}

	b.SetBotCommands(ctx, botCommands, gogram.WithLanguage("en"))

	b.RegisterHandler(gogram.HandlerTypeMessageText, "hello", gogram.MatchTypeCommandStartOnly, helloHandler)

	b.Start(ctx)
}

func helloHandler(ctx context.Context, bot *gogram.Bot, update *models.Update) {
	bot.SendMessageHelper(ctx, fmt.Sprintf("Hello, %s", update.Message.From.FirstName), update.Message.Chat.Id)
}

func printMessageMiddleware(next gogram.HandlerFunc) gogram.HandlerFunc {
	return func(ctx context.Context, bot *gogram.Bot, update *models.Update) {
		if update.Message != nil {
			fmt.Println("Message: ", update.Message.Text)
		}

		next(ctx, bot, update)
	}
}
