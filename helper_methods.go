package gogram

import (
	"context"

	"github.com/Alsond5/gogram/models"
)

func (b *Bot) SendMessageHelper(ctx context.Context, message string, chatId int64) (*models.Message, error) {
	m, err := b.SendMessage(ctx, &models.SendMessageRequest{
		ChatId: chatId,
		Text:   message,
	})

	return m, err
}
