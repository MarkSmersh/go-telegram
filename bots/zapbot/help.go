package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) help(e general.Message) {
	commands := ""

	for n, c := range b.Commands {
		commands += fmt.Sprintf("\n<code>%s</code> - <i>%s</i>", n, c.Description)
	}

	b.Tg.SendMessage(methods.SendMessage{
		Text:      fmt.Sprintf("Use <b>man</b> <u>command</u> to show a full description.\n%s", commands),
		ChatID:    e.Chat.ID,
		ParseMode: "HTML",
	})
}
