package zapbot

import (
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) rm(e general.Message) {
	if admin := b.isAdmin(e.Chat, *e.From); !admin {
		return
	}

	c, _ := core.StringToCommand(e.Text)

	if e.ReplyToMessage != nil {
		go b.Tg.DeleteMessage(methods.DeleteMessage{
			MessageID: e.ReplyToMessage.MessageID,
			ChatID:    e.Chat.ID,
		})
	}

	if t := c.GetArg("t"); t != "" {
		min := 1
		max := 100
		t, ok := b.optionAtoi(e, t, "t", &min, &max)

		if !ok {
			return
		}

		deleted := 0

		for i := 1; i <= 1000; i++ {
			if deleted >= t {
				break
			}

			_, err := b.Tg.DeleteMessage(methods.DeleteMessage{
				ChatID:    e.Chat.ID,
				MessageID: e.MessageID - i,
			})

			if err == nil {
				deleted++
			}
		}
	}

	if c.IsArg("s") {
		end := 0

		min := 1
		s, ok := b.optionAtoi(e, c.GetArg("s"), "-s", &min, nil)

		if !ok {
			return
		}

		if c.IsArg("e") {

			endOpt, ok := b.optionAtoi(e, c.GetArg("e"), "-e", &min, nil)

			if ok {
				end = endOpt
			}
		}

		if end != 0 && s >= end {
			b.Tg.SendMessage(methods.SendMessage{
				Text:      "The value of the option <code>-e</code> must be higher than the value of the option <code>-s</code>",
				ChatID:    e.Chat.ID,
				ParseMode: "HTML",
			})

			return
		}

		if end == 0 {
			end = e.MessageID - 1
		}

		for i := s; i <= end; i++ {
			go b.Tg.DeleteMessage(methods.DeleteMessage{
				ChatID:    e.Chat.ID,
				MessageID: i,
			})
		}

	}
	b.Tg.DeleteMessage(methods.DeleteMessage{
		MessageID: e.MessageID,
		ChatID:    e.Chat.ID,
	})
}
