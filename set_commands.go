package gogram

import (
	"context"
	"fmt"

	"github.com/Alsond5/gogram/models"
)

func (b *Bot) SetBotCommands(ctx context.Context, commands []models.BotCommand, options ...CommandOption) error {
	if len(commands) == 0 {
		return fmt.Errorf("commands list cannot be empty")
	}

	if len(commands) > 100 {
		return fmt.Errorf("maximum 100 commands allowed, got %d", len(commands))
	}

	config := &commandConfig{}

	for _, option := range options {
		option(config)
	}

	params := &models.SetMyCommandsRequest{
		Commands:     commands,
		LanguageCode: config.languageCode,
		Scope:        config.scope,
	}

	ok, err := b.SetMyCommands(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to set bot commands: %w", err)
	}

	if !ok {
		return fmt.Errorf("telegram API returned false for SetMyCommands")
	}

	return nil
}
