package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) man(e general.Message) {
	command, _ := core.StringToCommand(e.Text)

	if command.Etc == "" {
		b.Tg.SendMessage(methods.SendMessage{
			Text:      "What manual do you want? For example, try <code>man man</code> ",
			ChatID:    e.Chat.ID,
			ParseMode: "HTML",
		})

		return
	}

	c, ok := b.Commands[command.Etc]

	if !ok {
		b.Tg.SendMessage(methods.SendMessage{
			Text:      fmt.Sprintf("No manual entry for <code>%s</code>", command.Etc),
			ChatID:    e.Chat.ID,
			ParseMode: "HTML",
		})

		return
	}

	b.Tg.SendMessage(methods.SendMessage{
		ChatID:    e.Chat.ID,
		Text:      fmt.Sprintf("<code>%s</code> - <i>%s</i>\n\n%s", command.Etc, c.Description, c.Manual),
		ParseMode: "HTML",
	})
}
