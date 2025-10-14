package zapbot

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/redis/go-redis/v9"
)

func (b *ZapBot) export(e core.Message) {
	cli, _ := cli.NewCli(e.Raw().Text)
	splittedEtc := strings.Split(cli.Etc, " ")
	chatId := e.Chat.Raw().ID

	if cli.Exists("a") {
		variables, err := scanExportVariable(b.Rdb, chatId)

		if err != nil {
			fallbacks.UnableToProceed(e)
			return
		}

		if len(variables) <= 0 {
			e.Reply("You have no variables.")
			return
		}

		formattedVariables := []string{}

		for _, v := range variables {
			formattedVariables = append(
				formattedVariables,
				fmt.Sprintf("<code>%s</code>", v),
			)
		}

		e.Reply(strings.Join(formattedVariables, "\n"))
		return
	}

	deleteOption, _ := cli.Get("d")

	if deleteOption.Exists() {
		if deleteOption.IsEmpty() {
			fallbacks.OptionInvalidValue(e, deleteOption)
			return
		}

		variables := []string{deleteOption.Value}
		variablesFromEtc := strings.Split(cli.Etc, " ")
		variables = append(variables, variablesFromEtc...)

		deletedCount, err := deleteExportVariables(b.Rdb, chatId, variables...)

		if err != nil {
			slog.Error(err.Error())
			fallbacks.UnableToProceed(e)
			return
		}

		e.Reply(
			fmt.Sprintf("Succesfully deleted %d variables", deletedCount),
		)

		return
	}

	if e.ReplyToMessage != nil {
		if len(splittedEtc) <= 0 {
			fallbacks.InvalidFormat(e, cli, "Name for the variable is not given")
			return
		}

		if len(splittedEtc[0]) <= 0 {
			fallbacks.InvalidFormat(e, cli, "Variable's name length less than zero")
			return
		}

		err := writeExportVariable(b.Rdb, chatId, splittedEtc[0], e.ReplyToMessage.Raw().Text)

		if err != nil {
			slog.Error(err.Error())
			fallbacks.UnableToProceed(e)
			return
		}

		e.Reply(
			fmt.Sprintf("Varible <code>%s</code> is written", splittedEtc[0]),
		)

		return
	}

	if len(splittedEtc) == 1 {
		fallbacks.InvalidFormat(e, cli, "Value for the variable is not given.")
		return
	}

	if len(splittedEtc) <= 1 {
		fallbacks.InvalidFormat(e, cli, "")
		return
	}

	if len(splittedEtc[0]) <= 0 {
		fallbacks.InvalidFormat(e, cli, "Variable's name length less than zero")
		return
	}

	if len(splittedEtc[1]) <= 0 {
		fallbacks.InvalidFormat(e, cli, "Variable's value length less than zero")
		return
	}

	err := writeExportVariable(b.Rdb, chatId, splittedEtc[0], splittedEtc[1])

	if err != nil {
		slog.Error(err.Error())
		fallbacks.UnableToProceed(e)
		return
	}

	e.Reply(
		fmt.Sprintf("Varible <code>%s</code> is written", splittedEtc[0]),
	)
}

func writeExportVariable(rdb *redis.Client, chatId int, name string, value string) error {
	_, err := rdb.Set(
		context.Background(),
		fmt.Sprintf("export:%d:%s", chatId, name),
		value,
		0,
	).Result()
	return err
}

func scanExportVariable(rdb *redis.Client, chatId int) ([]string, error) {
	variables := []string{}
	formattedVariables := []string{}
	var cursor uint64 = 0

	newVariables, cursor, err := rdb.Scan(
		context.Background(),
		uint64(cursor),
		fmt.Sprintf("export:%d:*", chatId),
		1000,
	).Result()

	if err != nil {
		return variables, err
	}

	variables = append(variables, newVariables...)

	for cursor != 0 {
		newVariables, cursor, err = rdb.Scan(
			context.Background(),
			uint64(cursor),
			fmt.Sprintf("export:%d:*", chatId),
			1000,
		).Result()

		if err != nil {
			return variables, err
		}

		variables = append(variables, newVariables...)
	}

	for _, v := range variables {
		formattedVariable := strings.Split(v, ":")[2]
		formattedVariables = append(formattedVariables, formattedVariable)
	}

	return formattedVariables, nil
}

func deleteExportVariables(rdb *redis.Client, chatId int, names ...string) (int64, error) {
	rows := []string{}

	for _, name := range names {
		rows = append(rows, fmt.Sprintf("export:%d:%s", chatId, name))
	}

	return rdb.Del(context.Background(), rows...).Result()
}
