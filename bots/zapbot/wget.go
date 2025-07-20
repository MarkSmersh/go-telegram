package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) wget(e general.Message) {
	if e.ReplyToMessage == nil {
		b.Tg.SendMessage(methods.SendMessage{
			Text:   "Reply to message to get information about it.",
			ChatID: e.Chat.ID,
		})

		return
	}

	b.Tg.SendMessage(methods.SendMessage{
		Text:      fmt.Sprintf("Message ID: <code>%d</code>\nFrom ID: <code>%d</code>", e.ReplyToMessage.MessageID, e.ReplyToMessage.From.ID),
		ChatID:    e.Chat.ID,
		ParseMode: "HTML",
	})
}
