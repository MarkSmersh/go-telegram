package main

import (
	"log"
	"log/slog"
	"sync"

	"github.com/MarkSmersh/go-telegram/bots/zapbot"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
)

var env, _ = helpers.GetEnv()

var bot = zapbot.ZapBot{
	Tg: core.Telegram{
		Token:   env["BOT_TOKEN_2"],
		Eventer: core.Updater{},
	},
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	go bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
