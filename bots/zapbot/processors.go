package zapbot

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) ProcessInlines(e general.InlineQuery) {

}

func (b *ZapBot) ProcessCallbacks(e general.CallbackQuery) {
	slog.Debug(
		fmt.Sprintf("data: %v", e.Data),
	)

	splittedData := strings.Split(e.Data, ":")

	if len(splittedData) < 2 {
		return
	}

	slog.Debug(
		fmt.Sprintf("splittedData: %v", splittedData),
	)

	if splittedData[0] == "write" {
		var data WriteData
		err := b.Rdb.HGetAll(context.Background(), e.Data).Scan(&data)

		if err != nil {
			slog.Error(
				err.Error(),
			)
			return
		}

		slog.Debug("write", "eFromId", e.From.ID, "dataTarget", data.Target)

		if e.From.ID != data.Target {
			b.Tg.AnswerCallbackQuery(methods.AnswerCallbackQuery{
				CallbackQueryID: e.ID,
				Text:            "This is not your domain, my dear.",
			})

			return
		}

		b.Tg.AnswerCallbackQuery(methods.AnswerCallbackQuery{
			CallbackQueryID: e.ID,
			Text:            data.Data,
			ShowAlert:       true,
		})
	}
}

func (b *ZapBot) ProcessCommands(e general.Message) {
	go b.RegisterUsername(e.From.ID, e.From.Username)

	if e.Text == "" {
		return
	}

	cli, ok := cli.NewCli(e.Text)

	if !ok {
		return
	}

	slog.Debug(
		fmt.Sprintf("Parsed command:\nCommand: %s\nargs: %v\netc: %s", cli.Entry, cli.Args, cli.Etc),
	)

	command, ok := b.CommandsMng.Get(cli.Entry)

	if !ok {
		return
	}

	message := b.Tg.NewMessage(e)

	// let's set it as a standart one for the entire zapbot
	message.SetParseMode(consts.HTML)

	for _, middleware := range command.Middlewares {
		if !middleware(message) {
			return
		}
	}

	command.Function(message)
}
