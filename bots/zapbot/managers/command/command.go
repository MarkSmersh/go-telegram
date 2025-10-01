package command

import "github.com/MarkSmersh/go-telegram/core"

type Command struct {
	Description string
	Manual      string
	Middlewares []func(core.Message) bool
	Function    func(core.Message)
}
