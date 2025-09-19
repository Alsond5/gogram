package gogram

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
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
	handlers           []handler
	handlersMx         sync.RWMutex
	middlewares        []Handler
	middlewaresMx      sync.RWMutex
	middlewareCount    int64
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

func (b *Bot) Use(handler Handler) {
	b.middlewaresMx.Lock()
	defer b.middlewaresMx.Unlock()

	b.middlewares = append(b.middlewares, handler)
	atomic.AddInt64(&b.middlewareCount, 1)
}

func (b *Bot) getMiddlewareCount() int {
	return int(atomic.LoadInt64(&b.middlewareCount))
}

func (b *Bot) getMiddlewares(count int) []Handler {
	if count < 0 {
		return nil
	}

	b.middlewaresMx.RLock()
	defer b.middlewaresMx.RUnlock()

	total := len(b.middlewares)
	if total == 0 {
		return nil
	}

	if count > total {
		count = total
	}

	result := make([]Handler, count)
	copy(result, b.middlewares[:count])

	return result
}
