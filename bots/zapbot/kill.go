package zapbot

import (
	"fmt"
	"strconv"
	"time"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) kill(e general.Message) {
	if !b.isAdmin(e.Chat, *e.From) {
		return
	}

	c, _ := core.StringToCommand(e.Text)

	untilDate := time.Now().UnixMilli() + 60

	target := 0

	if e.ReplyToMessage != nil {
		target = e.ReplyToMessage.From.ID
	}

	if c.IsArg("f") || c.IsArg("ban") {
		untilDate = 0
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
				Text:   "Value for an option --user-id is incorrect",
			})
			return
		}

		target = t
	}

	if target <= 0 {
		b.Tg.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   "Option User ID is missing. Reply to message's owner you want to be banned. Also you can use -u or --user-id to set user'd is manually.",
		})

		return
	}

	m, _ := b.Tg.GetChatMember(methods.GetChatMember{
		ChatID: e.Chat.ID,
		UserID: target,
	})

	if m.Status == consts.CREATOR {
		b.Tg.SendMessage(methods.SendMessage{
			Text:   "You are cooked.",
			ChatID: e.Chat.ID,
		})
		return
	}

	if m.Status == consts.ADMINISTRATOR {
		b.Tg.SendMessage(methods.SendMessage{
			Text:   "Admin cannot be killed.",
			ChatID: e.Chat.ID,
		})
		return
	}

	_, err := b.Tg.BanChatMember(methods.BanChatMember{
		ChatID:    e.Chat.ID,
		UserID:    target,
		UntilDate: int(untilDate),
	})

	if err != nil {
		b.Tg.SendMessage(methods.SendMessage{
			Text:   fmt.Sprintf("User cannot be removed.\n\n%s", err.Error()),
			ChatID: e.Chat.ID,
		})
		return
	}

	if untilDate == 0 {
		b.Tg.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   fmt.Sprintf("- Well... done... %s...\n- Farewell, %s.", e.From.FirstName, m.User.FirstName),
		})
	} else {
		b.Tg.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   fmt.Sprintf("Killed %s.", m.User.FirstName),
		})
	}
}
