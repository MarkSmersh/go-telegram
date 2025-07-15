package anonchat

import (
	"strconv"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/core/keyboard"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) interests(e general.Message) {
	if e.Text == "/interests" {
		u := b.Chat.Users[e.From.ID].Interests

		keyboard := keyboard.ReplyMarkup{
			InlineButtons: [][]keyboard.InlineButton{
				{
					keyboard.InlineButton{
						Text:         helpers.InterestToStr(consts.InterestTalking) + " " + helpers.IsInterestIn(u, consts.InterestTalking),
						CallbackData: "i-" + strconv.Itoa(consts.InterestTalking),
					},
					keyboard.InlineButton{
						Text:         helpers.InterestToStr(consts.InterestSex) + " " + helpers.IsInterestIn(u, consts.InterestSex),
						CallbackData: "i-" + strconv.Itoa(consts.InterestSex),
					},
				},
			},
		}

		b.Tg.SendMessage(methods.SendMessage{
			ChatID:      e.Chat.ID,
			Text:        "Choose your interesets from below:",
			ReplyMarkup: keyboard.ToJSON(),
		})
	}
}
