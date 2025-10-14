package zapbot

import (
	"context"
	"fmt"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/redis/go-redis/v9"
)

func (b *ZapBot) echo(e core.Message) {
	cli, _ := cli.NewCli(e.Raw().Text)

	if len(cli.Etc) <= 0 {
		fallbacks.InvalidFormat(e, cli, "Variable name is not provided")
		return
	}

	value, err := getExportVariable(b.Rdb, e.Chat.Raw().ID, cli.Etc)

	if err == nil && len(value) <= 0 {
		e.Reply(
			fmt.Sprintf("Variable %s doesn't exist", cli.Etc),
		)
		return
	}

	if err != nil {
		fallbacks.UnableToProceed(e)
		return
	}

	e.DisableLinkPreview()
	e.Reply(value)
}

func getExportVariable(rdb *redis.Client, chatId int, name string) (string, error) {
	return rdb.Get(
		context.Background(),
		fmt.Sprintf("export:%d:%s", chatId, name),
	).Result()
}
