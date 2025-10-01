package zapbot

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-telegram/core"
)

func (b *ZapBot) ping(m core.Message) {
	t := time.Now().UnixMilli()

	b.Tg.GetMe()
	m.Reply(fmt.Sprintf("Pong: %dms", time.Now().UnixMilli()-t))
}
