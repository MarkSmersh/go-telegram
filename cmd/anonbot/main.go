package main

import (
	"sync"

	"github.com/MarkSmersh/go-telegram/bots/anonchat"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
)

var env, _ = helpers.GetEnv()

var bot = anonchat.AnonBot{
	Tg: core.Telegram{
		Token:   env["BOT_TOKEN_1"],
		Eventer: core.Updater{},
	},
	Chat:  core.Chat{},
	Users: core.State[int, int]{},
}

func main() {
	go bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
