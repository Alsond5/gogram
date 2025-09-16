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

	b, err := gogram.NewBot(token, gogram.SetWorkers(3))
	if err != nil {
		fmt.Println(err)
	}

	b.SetBotCommands(ctx, botCommands, gogram.WithLanguage("en"))

	b.Use(printMessageMiddleware())

	b.RegisterHandler(gogram.HandlerTypeMessageText, "hello", gogram.MatchTypeCommandStartOnly, helloHandler)

	b.Start(ctx)
}

func helloHandler(c *gogram.Context) error {
	_, err := c.Bot.SendMessageHelper(c.Ctx, fmt.Sprintf("Hello, %s", c.Update.Message.From.FirstName), c.Update.Message.Chat.Id)

	return err
}

func printMessageMiddleware() gogram.Handler {
	return func(c *gogram.Context) error {
		update := c.Update

		if update.Message != nil {
			fmt.Println("Message: ", update.Message.Text)
		}

		return c.Next()
	}
}
