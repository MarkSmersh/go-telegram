package anonchat

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) ping(e general.Message) {
	if e.Text == "/ping" {
		before := time.Now().UnixMilli()

		m, _ := b.Tg.SendMessage(methods.SendMessage{
			Text:   "Ping...",
			ChatID: e.Chat.ID,
		})

		after := time.Now().UnixMilli()

		b.Tg.EditMessageText(methods.EditMessageText{
			Text:      fmt.Sprintf("Pong: %d ms", after-before),
			MessageID: m.MessageID,
			ChatID:    m.Chat.ID,
		})
	}
}
