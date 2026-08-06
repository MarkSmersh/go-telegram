package klipy

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MarkSmersh/go-telegram/components/klipy"
	"github.com/MarkSmersh/go-telegram/components/klipy/types/methods"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
)

type KlipyBot struct {
	tg core.Telegram
}

func NewKlipyBot(token string) KlipyBot {
	return KlipyBot{
		core.NewTelegram(token),
	}
}

func (k KlipyBot) Init() {

	klipy := klipy.NewKlipy(os.Getenv("KLIPY_API"))

	k.tg.Eventer.InlineQuery.Add(func(iq general.InlineQuery) {
		slog.Debug(iq.Query)

		res, err := klipy.Search(methods.Search{
			Q: iq.Query,
		})

		if err != nil {
			slog.Error(err.Error())
			return
		}

		slog.Debug(fmt.Sprintf("RESPONSE: %d", len(res)))

		ib := core.NewInlineBuilder(iq.ID)

		for _, r := range res {
			ib.AddGIF(core.InlineQueryResultMpeg4Gif{
				Type:         "mpeg4_gif",
				ID:           fmt.Sprintf("%d", r.ID),
				Mpeg4URL:     r.File.MD.GIF.URL,
				ThumbnailURL: r.File.XS.GIF.URL,
			})
		}

		_, err = k.tg.AnswerInlineQuery(ib.AnswerInlineQuery)

		if err != nil {
			slog.Error(err.Error())
		}
	})

	k.tg.Eventer.Messages.Add(func(m general.Message) {
		k.tg.NewMessage(m).Reply(m.Text)
	})

	k.tg.Init(func(e core.User) {
		slog.Info(
			fmt.Sprintf("%s has started", e.Raw().FirstName),
		)
	})
}
