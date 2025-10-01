package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/core"
)

func (b *ZapBot) help(e core.Message) {
	commands := ""

	for n, c := range b.CommandsMng.Commands() {
		commands += fmt.Sprintf("\n<code>%s</code> - <i>%s</i>", n, c.Description)
	}

	e.Reply(fmt.Sprintf("Use <b>man</b> <u>command</u> to show a full description.\n%s", commands))
}
