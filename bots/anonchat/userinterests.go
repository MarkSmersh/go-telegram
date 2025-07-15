package anonchat

import (
	"strconv"
	"strings"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/core/keyboard"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *AnonBot) userInterests(e general.CallbackQuery) {
	splittedQuery := strings.Split(e.Data, "-")

	prefix := splittedQuery[0]
	value := splittedQuery[1]

	if prefix != "i" {
		return
	}

	v, err := strconv.Atoi(value)

	if err != nil {
		return
	}

	b.Chat.Users[e.From.ID].AddOrRemoveInterest(v)

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

	b.Tg.EditMessageReplyMarkup(methods.EditMessageReplyMarkup{
		ReplyMarkup: keyboard.ToJSON(),
		MessageID:   e.Message.MessageID,
		ChatID:      e.From.ID,
	})
}
