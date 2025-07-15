package anonchat

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) start(e general.Message) {
	if e.Text == "/start" {
		b.Tg.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   fmt.Sprintf("Hello, %s. Welcome to the edging zone.", e.From.FirstName),
		})
	}
}
