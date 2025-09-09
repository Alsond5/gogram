package gogram

type Middleware func(next HandlerFunc) HandlerFunc
