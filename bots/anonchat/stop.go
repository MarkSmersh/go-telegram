package anonchat

import (
	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) stop(e general.Message) {
	if e.Text == "/stop" {
		if b.Users.Get(e.Chat.ID) == consts.StateDefault {
			b.Tg.SendMessage(methods.SendMessage{
				Text:   "You have no companion. You are not even searching for him. What are trying to do?",
				ChatID: e.Chat.ID,
			})
			return
		}

		b.Chat.RemoveFromSearch(e.Chat.ID)
		co := b.Chat.Disconnect(e.Chat.ID)
		b.Users.Set(e.Chat.ID, consts.StateDefault)

		if co != 0 {
			b.Users.Set(co, consts.StateDefault)

			b.Tg.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "You've stopped the dialogue. Use /next to search a new one",
			})

			b.Tg.SendMessage(methods.SendMessage{
				ChatID: co,
				Text:   "You've been skipped by your companion. There is no need to worry. You can find another one with /next",
			})
		} else {
			b.Tg.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "You've stopped searching process",
			})
		}
	}
}
