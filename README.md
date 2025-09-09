# Gogram (Telegram Bot Framework)

Gogram is a **Go framework for building Telegram bots**. It simplifies bot development with features like middleware support, command management, inline buttons, and callback handling.

## Features

- **Easy bot creation:** Initialize your bot with a single `NewBot` call.  
- **Middleware support:** Preprocess messages, log activity, or apply filters easily.  
- **Command management:** Add bot commands such as `/start`, `/help`, `/hello` with ease.  
- **Worker management:** Adjust the number of concurrent workers to optimize performance.  
- **Inline buttons and callback handling:** Build interactive bots with Telegram inline keyboards.

## Installation

```bash
go get github.com/Alsond5/gogram
```

## Example Usage

```go
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
```

## License

Gogram is licensed under the **Apache License 2.0**. See the [LICENSE](LICENSE) file for details.