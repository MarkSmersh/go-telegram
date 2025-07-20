package zapbot

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) ping(e general.Message) {
	t := time.Now().UnixMilli()

	m, _ := b.Tg.SendMessage(methods.SendMessage{
		Text:   "Ping...",
		ChatID: e.Chat.ID,
	})

	b.Tg.EditMessageText(methods.EditMessageText{
		Text:      fmt.Sprintf("Pong: %dms", time.Now().UnixMilli()-t),
		MessageID: m.MessageID,
		ChatID:    e.Chat.ID,
	})
}
