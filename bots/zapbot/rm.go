package zapbot

import (
	"encoding/json"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) rm(e core.Message) {
	c, _ := cli.NewCli(e.Raw().Text)

	if e.ReplyToMessage != nil {
		go e.ReplyToMessage.DeleteMessage()
	}

	tailOption, ok := c.Get("t")

	if ok && !tailOption.IsEmpty() {
		tail, ok := tailOption.AtoiRange(1, 100)

		if !ok {
			fallbacks.OptionInvalidValue(e, tailOption)
			return
		}

		deletedCount := 0

		for i := 1; i <= 1000; i++ {
			if deletedCount >= tail {
				break
			}

			// _, err := e.DeleteMessage()

			_, err := b.Tg.DeleteMessage(methods.DeleteMessage{
				ChatID:    e.Raw().Chat.ID,
				MessageID: e.Raw().MessageID - i,
			})

			if err == nil {
				deletedCount++
			}
		}
	}

	startOption, ok := c.Get("s")

	if ok {
		start, ok := startOption.Atoi()
		end := 0

		if !ok || start < 1 {
			fallbacks.OptionInvalidValue(e, startOption)
			return
		}

		endOption, ok := c.Get("e")

		if ok {
			end, ok = endOption.Atoi()

			if !ok {
				fallbacks.OptionInvalidValue(e, endOption)
				return
			}

			if end != 0 && end <= start {
				e.Reply("The value of the option <code>-e</code> must be higher than the value of the option <code>-s</code>")
				return
			}
		}

		if end == 0 {
			end = e.Raw().MessageID - 1
		}

		// TODO: end
		for i := start; i <= end; i = i + 100 {
			messages := []int{}

			for j := range end - i + 1 {
				messages = append(messages, i+j)
			}

			jsonMessages, _ := json.Marshal(messages)

			go b.Tg.DeleteMessages(methods.DeleteMessages{
				ChatID:     e.Chat.Raw().ID,
				MessageIDs: string(jsonMessages),
			})
		}
	}

	e.DeleteMessage()
}
