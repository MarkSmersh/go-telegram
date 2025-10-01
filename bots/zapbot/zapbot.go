package zapbot

import (
	"github.com/MarkSmersh/go-telegram/bots/zapbot/managers/command"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/redis/go-redis/v9"
)

type ZapBot struct {
	Tg          core.Telegram
	CommandsMng command.CommandsManager
	Rdb         *redis.Client
}

func NewZapBot(tg core.Telegram, rdb *redis.Client) ZapBot {
	return ZapBot{
		Tg:          tg,
		Rdb:         rdb,
		CommandsMng: command.NewCommandsManager(),
	}
}

func (b *ZapBot) Init() {
	b.CheckRedis()
	b.SetUpEventers()
	b.RegisterCommands()

	b.Tg.Init(b.onInit)
}
