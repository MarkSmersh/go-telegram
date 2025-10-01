package zapbot

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/core/keyboard"
	"github.com/MarkSmersh/go-telegram/types/methods"
	"github.com/google/uuid"
)

type WriteData struct {
	Data   string `redis:"data"`
	Target int    `redis:"target"`
}

func (b *ZapBot) write(e core.Message) {
	if e.Chat.Raw().ID != e.From.Raw().ID {
		go e.DeleteMessage()
	}

	cli, _ := cli.NewCli(e.Raw().Text)

	chatId := e.Chat.Raw().ID
	target := 0

	if e.ReplyToMessage != nil {
		target = e.ReplyToMessage.From.Raw().ID
	}

	if target == 0 {
		user, ok := cli.Get("u")

		if !ok {
			return
		}

		id, err := b.ExtractUserId(user)

		if err != nil {
			fallbacks.OptionInvalidValue(e, user)
			return
		}

		target = id
	}

	if target == 0 {
		b.Tg.SendMessage(methods.SendMessage{
			Text:   "No user provided.",
			ChatID: e.Chat.Raw().ID,
		})

		return
	}

	id := uuid.New().String()

	data := []string{
		"data", cli.Etc,
		"target", strconv.Itoa(target),
	}

	_, err := b.Rdb.HSet(context.Background(), "write:"+id, data).Result()

	if err != nil {
		slog.Error(
			err.Error(),
		)
	}

	kb := keyboard.ReplyMarkup{
		InlineButtons: [][]keyboard.InlineButton{
			{
				{
					Text:         "See the content",
					CallbackData: "write:" + id,
				},
			},
		},
	}

	_, err = b.Tg.SendMessage(methods.SendMessage{
		Text: fmt.Sprintf(
			"%s, %s wrote you something\\. Check it out\\!",
			b.GetUserMention(target),
			b.GetUserMention(e.From.Raw().ID),
		),
		ParseMode:   "MarkdownV2",
		ReplyMarkup: kb.ToJSON(),
		ChatID:      chatId,
	})

	if err != nil {
		fallbacks.ApiError(e, err.Error(), "Message cannot be sent.")
	}
}
