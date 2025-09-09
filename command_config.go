package gogram

import "github.com/Alsond5/gogram/models"

type CommandOption func(*commandConfig)

type commandConfig struct {
	languageCode string
	scope        *models.BotCommandScope
}

func WithLanguage(languageCode string) CommandOption {
	return func(cc *commandConfig) {
		cc.languageCode = languageCode
	}
}

func WithScope(scope *models.BotCommandScope) CommandOption {
	return func(cc *commandConfig) {
		cc.scope = scope
	}
}
