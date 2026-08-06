package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/MarkSmersh/go-telegram/bots/klipy"
	"github.com/MarkSmersh/go-telegram/helpers"
)

func main() {
	helpers.LoadDotEnv(".env")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	bot := klipy.NewKlipyBot(os.Getenv("KLIPY_TOKEN"))

	go bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
