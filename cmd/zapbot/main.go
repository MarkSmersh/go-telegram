package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/MarkSmersh/go-telegram/bots/zapbot"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/redis/go-redis/v9"
)

func main() {
	helpers.LoadDotEnv(".env")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	var rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ZAPBOT_REDIS_ADDR"),
		Password: os.Getenv("ZAPBOT_REDIS_PASSWORD"),
		DB:       helpers.GetEnvInt("ZAPBOT_REDIS_DB"),
	})

	var tg = core.NewTelegram(
		os.Getenv("ZAPBOT_TOKEN"),
	)

	var bot = zapbot.NewZapBot(tg, rdb)

	go bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
