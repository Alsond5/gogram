package gogram

import (
	"context"
	"regexp"
	"strings"

	"github.com/Alsond5/gogram/models"
)

type HandlerType int

const (
	HandlerTypeMessageText HandlerType = iota
	HandlerTypeCallbackQueryData
	HandlerTypeCallbackQueryGameShortName
	HandlerTypePhotoCaption
)

type MatchType int

const (
	MatchTypeExact MatchType = iota
	MatchTypePrefix
	MatchTypeContains
	MatchTypeCommand
	MatchTypeCommandStartOnly
	matchTypeRegexp
	matchTypeFunc
)

type HandlerFunc func(ctx context.Context, bot *Bot, update *models.Update)
type MatchFunc func(upd *models.Update) bool

type handler struct {
	handlerType HandlerType
	matchType   MatchType
	handler     HandlerFunc

	pattern   string
	re        *regexp.Regexp
	matchFunc MatchFunc
}

func (h *handler) match(upd *models.Update) bool {
	if h.matchType == matchTypeFunc {
		return h.matchFunc(upd)
	}

	var data string
	var entities []models.MessageEntity

	switch h.handlerType {
	case HandlerTypeMessageText:
		if upd.Message == nil {
			return false
		}

		data = upd.Message.Text
		entities = upd.Message.Entities
	case HandlerTypeCallbackQueryData:
		if upd.CallbackQuery == nil {
			return false
		}

		data = upd.CallbackQuery.Data
	case HandlerTypeCallbackQueryGameShortName:
		if upd.CallbackQuery == nil {
			return false
		}

		data = upd.CallbackQuery.GameShortName
	case HandlerTypePhotoCaption:
		if upd.Message == nil {
			return false
		}

		data = upd.Message.Caption
		entities = upd.Message.CaptionEntities
	}

	if h.matchType == MatchTypeExact {
		return data == h.pattern
	}

	if h.matchType == MatchTypePrefix {
		return strings.HasPrefix(data, h.pattern)
	}

	if h.matchType == MatchTypeContains {
		return strings.Contains(data, h.pattern)
	}

	if h.matchType == MatchTypeCommand {
		for _, e := range entities {
			if e.Type != "bot_command" {
				continue
			}

			if data[e.Offset+1:e.Offset+e.Length] == h.pattern {
				return true
			}
		}
	}

	if h.matchType == MatchTypeCommandStartOnly {
		for _, e := range entities {
			if e.Type != "bot_command" {
				continue
			}

			if e.Offset == 0 && data[e.Offset+1:e.Offset+e.Length] == h.pattern {
				return true
			}
		}
	}

	if h.matchType == matchTypeRegexp {
		return h.re.Match([]byte(data))
	}

	return false
}

func (b *Bot) OnStartCommand(handlerFunc HandlerFunc) {
	b.RegisterCommandHandler("start", handlerFunc)
}

func (b *Bot) OnHelpCommand(handlerFunc HandlerFunc) {
	b.RegisterCommandHandler("help", handlerFunc)
}

func (b *Bot) RegisterCommandHandler(command string, handlerFunc HandlerFunc) {
	b.RegisterHandler(HandlerTypeMessageText, command, MatchTypeCommandStartOnly, handlerFunc)
}

func (b *Bot) RegisterHandler(handlerType HandlerType, pattern string, matchType MatchType, handlerFunc HandlerFunc) {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	h := handler{
		pattern:     pattern,
		handlerType: handlerType,
		matchType:   matchType,
		handler:     handlerFunc,
	}

	b.handlers = append(b.handlers, h)
}
