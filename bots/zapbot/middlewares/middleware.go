package middlewares

import "github.com/MarkSmersh/go-telegram/core"

type Middleware struct {
	tg core.Telegram
}

func NewMiddleware(tg core.Telegram) Middleware {
	return Middleware{
		tg: tg,
	}
}
