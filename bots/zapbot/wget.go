package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/core"
)

func (b *ZapBot) wget(e core.Message) {
	if e.ReplyToMessage == nil {
		e.Reply("Reply to the message to get information about it.")
		return
	}

	e.Reply(
		fmt.Sprintf(
			"Message ID: <code>%d</code>\nChat ID: <code>%d</code>:\nFrom ID: <code>%d</code>\nUsername: @%s",
			e.ReplyToMessage.Raw().MessageID,
			e.Chat.Raw().ID,
			e.ReplyToMessage.Raw().From.ID,
			b.GetUsername(e.ReplyToMessage.Raw().From.ID)),
	)
}
