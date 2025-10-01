package zapbot

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/managers/command"
	"github.com/MarkSmersh/go-telegram/bots/zapbot/middlewares"
	"github.com/MarkSmersh/go-telegram/core"
)

func (b ZapBot) CheckRedis() {
	_, err := b.Rdb.Ping(context.Background()).Result()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func (b *ZapBot) SetUpEventers() {
	b.Tg.Eventer.Messages.Add(b.ProcessCommands)
	b.Tg.Eventer.CallbackQuery.Add(b.ProcessCallbacks)
	// b.Tg.Eventer.InlineQuery.Add(b.ProcessInline)
}

func (b *ZapBot) RegisterCommands() {
	manuals := command.ReadManualsFrom("./bots/zapbot/man")

	middleware := middlewares.NewMiddleware(b.Tg)

	b.CommandsMng.RegisterCommand(
		"man",
		"for manual",
		b.man,
		manuals["man"],
	)

	b.CommandsMng.RegisterCommand(
		"ping",
		"shows ping between host and Telegram servers",
		b.ping,
		manuals["ping"],
	)

	b.CommandsMng.RegisterCommand(
		"help",
		"shows a list of actual commands",
		b.help,
		manuals["help"],
	)

	b.CommandsMng.RegisterCommand(
		"kill",
		"removes user from the chat",
		b.kill,
		manuals["kill"],
		middleware.IsAdmin,
	)

	b.CommandsMng.RegisterCommand(
		"respawn",
		"respawns the user",
		b.respawn,
		manuals["respawn"],
		middleware.IsAdmin,
	)

	b.CommandsMng.RegisterCommand(
		"rm",
		"removes message/messages",
		b.rm,
		manuals["rm"],
		middleware.IsAdmin,
	)

	b.CommandsMng.RegisterCommand(
		"wget",
		"displays info about the replied message",
		b.wget,
		manuals["wget"],
	)

	b.CommandsMng.RegisterCommand(
		"write",
		"creates a message, content of which only provided user can see",
		b.write,
		manuals["write"],
	)

	b.CommandsMng.RegisterCommand(
		"rule34",
		"finds a superiour content by tags (beta)",
		b.rule34,
		manuals["rule34"],
	)
}

func (b *ZapBot) onInit(e core.User) {
	slog.Info(
		fmt.Sprintf("Bot %s started", e.Raw().FirstName),
	)
}
