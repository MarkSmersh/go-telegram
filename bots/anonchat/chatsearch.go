package anonchat

import (
	"fmt"
	"strings"
	"time"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) chatSearch(e general.Message) {
	if e.Text == "/next" {
		if b.Users.Get(e.Chat.ID) == consts.StateSearch {
			b.Tg.SendMessage(methods.SendMessage{
				Text:   "You are already searching for a companion. There is no need to spam it.",
				ChatID: e.Chat.ID,
			})
			return
		}

		if b.Users.Get(e.Chat.ID) == consts.StateConnected {
			b.Tg.SendMessage(methods.SendMessage{
				Text:   "You already have a companion. If you don't like use /stop and then /next. (it will be changed quite soon)",
				ChatID: e.Chat.ID,
			})
			return
		}

		b.Users.Set(e.Chat.ID, consts.StateSearch)

		text := "Searching...\n\nYou can stop a searching process with /stop"

		if len(b.Chat.Users[e.Chat.ID].Interests) > 0 {
			interests := []string{}

			for _, i := range b.Chat.Users[e.Chat.ID].Interests {
				interests = append(interests, helpers.InterestToStr(i))
			}

			text = fmt.Sprintf("Your interests: %s\n\n", strings.Join(interests, ", ")) + text
		}

		b.Tg.SendMessage(methods.SendMessage{
			Text:   text,
			ChatID: e.Chat.ID,
		})

		b.Chat.AddToSearch(e.Chat.ID)

		go func() {
			for {
				if b.Users.Get(e.Chat.ID) != consts.StateSearch {
					return
				}

				userId, equalInterests := b.Chat.GetFirstCompanion(e.Chat.ID)

				interests := []string{}

				for _, i := range equalInterests {
					interests = append(interests, helpers.InterestToStr(i))
				}

				if userId != 0 {
					b.Chat.Connect(e.Chat.ID, userId)

					b.Users.Set(e.Chat.ID, consts.StateConnected)
					b.Users.Set(userId, consts.StateConnected)

					text := "New companion is found (id%d)"

					if len(equalInterests) > 0 {
						text = fmt.Sprintf("Equal interests: %s\n\n", strings.Join(interests, ", ")) + text
					}

					b.Tg.SendMessage(methods.SendMessage{
						ChatID: e.Chat.ID,
						Text:   fmt.Sprintf(text, userId),
					})

					b.Tg.SendMessage(methods.SendMessage{
						ChatID: userId,
						Text:   fmt.Sprintf(text, e.Chat.ID),
					})

					return
				} else {
					time.Sleep(1000 * time.Millisecond)
				}
			}
		}()
	}
}
