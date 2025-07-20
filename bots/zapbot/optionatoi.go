package zapbot

import (
	"fmt"
	"strconv"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) optionAtoi(e general.Message, value string, option string, min *int, max *int) (int, bool) {
	v, err := strconv.Atoi(value)
	cooked := false

	if err != nil {
		cooked = true
	}

	if min != nil && v < *min {
		cooked = true
	}

	if max != nil && v > *max {
		cooked = true
	}

	if cooked {
		b.Tg.SendMessage(methods.SendMessage{
			Text:      fmt.Sprintf("Value for an option <code>%s</code> is incorrect.", option),
			ChatID:    e.Chat.ID,
			ParseMode: "HTML",
		})

		return 0, false
	}

	return v, true
}
