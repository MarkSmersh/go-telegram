package klipy

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

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
		page, _ := strconv.Atoi(iq.Offset)

		res, err := klipy.Search(methods.Search{
			Q:       iq.Query,
			PerPage: 50,
			Page:    page,
		})

		if err != nil {
			slog.Error(err.Error())
			return
		}

		ib := core.NewInlineBuilder(iq.ID)
		ib.AnswerInlineQuery.CacheTime = 60

		ib.AnswerInlineQuery.NextOffset = fmt.Sprintf("%d", res.CurrentPage+1)

		for _, r := range res.Data {
			ib.AddGIF(core.InlineQueryResultGif{
				Type:              "gif",
				ID:                fmt.Sprintf("%d", r.ID),
				GifURL:            r.File.HD.MP4.URL,
				ThumbnailURL:      r.File.XS.JPG.URL,
				Title:             r.Title,
				ThumbnailMimeType: "image/jpeg",
				GifWidth:          r.File.XS.GIF.Width,
				GifHeight:         r.File.XS.GIF.Height,
			})
		}

		_, err = k.tg.AnswerInlineQuery(ib.AnswerInlineQuery)

		if err != nil {
			slog.Error(err.Error())
		}
	})

	k.tg.Init(func(e core.User) {
		slog.Info(
			fmt.Sprintf("%s has started", e.Raw().FirstName),
		)
	})
}
