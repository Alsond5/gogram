package gogram

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Alsond5/gogram/models"
)

const (
	defaultPollTimeout    = time.Minute
	defaultUpdatesChanCap = 1024
	defaultRequestTimeout = time.Second * 5
	defaultWorkers        = 1
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Bot struct {
	url                string
	token              string
	workers            int
	pollTimeout        time.Duration
	requestTimeout     time.Duration
	webhookSecretToken string
	lastUpdateID       int64
	handlersMx         sync.RWMutex
	handlers           []handler
	middlewares        []Middleware
	client             HttpClient
	updates            chan *models.Update
}

func NewBot(token string, options ...Option) (*Bot, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("empty token")
	}

	bot := &Bot{
		url:            "https://api.telegram.org",
		token:          token,
		pollTimeout:    defaultPollTimeout,
		requestTimeout: defaultRequestTimeout,
		workers:        defaultWorkers,
		client: &http.Client{
			Timeout: defaultPollTimeout,
		},
		updates: make(chan *models.Update, defaultUpdatesChanCap),
	}

	for _, option := range options {
		option(bot)
	}

	ctx, cancel := context.WithTimeout(context.Background(), bot.requestTimeout)
	defer cancel()

	user, err := bot.GetMe(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while calling getMe: %w", err)
	}

	fmt.Printf("🤖 %s is up and running...\n", user.FirstName)

	return bot, nil
}

func (b *Bot) Start(ctx context.Context) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go b.startPooling(ctx, &wg)

	wg.Add(b.workers)
	for i := 0; i < b.workers; i++ {
		go b.handleUpdates(ctx, &wg)
	}

	wg.Wait()
}
