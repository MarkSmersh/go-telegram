package zapbot

import (
	"fmt"
	"strconv"

	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) respawn(e general.Message) {
	if !b.isAdmin(e.Chat, *e.From) {
		return
	}

	c, _ := core.StringToCommand(e.Text)

	target := 0

	if e.ReplyToMessage != nil {
		target = e.ReplyToMessage.From.ID
	}

	if u := c.GetArg("u"); u != "" {
		t, err := strconv.Atoi(u)

		if err != nil {
			b.Tg.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "Value for an option -u is incorrect.",
			})
			return
		}

		target = t
	}

	if u := c.GetArg("user-id"); u != "" {
		t, err := strconv.Atoi(u)

		if err != nil {
			b.Tg.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "Value for an option --user-id is incorrect.",
			})
			return
		}

		target = t
	}

	_, err := b.Tg.UnbanChatMember(methods.UnbanChatMember{
		ChatID: e.Chat.ID,
		UserID: target,
	})

	if err != nil {
		b.Tg.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   fmt.Sprintf("User cannot be respawned.\n\n%s", err.Error()),
		})
		return
	}

	m, _ := b.Tg.GetChatMember(methods.GetChatMember{
		ChatID: e.Chat.ID,
		UserID: target,
	})

	b.Tg.SendMessage(methods.SendMessage{
		ChatID: e.Chat.ID,
		Text:   fmt.Sprintf("Respawned %s. Invite him someone", m.User.FirstName),
	})
}
