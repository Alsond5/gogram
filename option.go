package gogram

type Option func(b *Bot)

func SetWorkers(count int) Option {
	return func(b *Bot) {
		b.workers = count
	}
}

func SetMiddlewares(middlewares ...Middleware) Option {
	return func(b *Bot) {
		b.middlewares = append(b.middlewares, middlewares...)
	}
}
