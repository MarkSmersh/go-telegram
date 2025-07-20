package zapbot

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/general"
)

type Command struct {
	Description string
	Manual      string
	Function    func(general.Message)
}

type ZapBot struct {
	Tg       core.Telegram
	Commands map[string]Command
}

func (b *ZapBot) Init() {
	b.Commands = map[string]Command{}

	b.Tg.Eventer.Messages.Add(b.ProcessCommands)

	manuals := b.ReadManualsFrom("/man")

	b.RegisterCommand(
		"man",
		"for manual",
		b.man,
		manuals["man"],
	)

	b.RegisterCommand(
		"ping",
		"shows ping between host and Telegram servers",
		b.ping,
		manuals["ping"],
	)

	b.RegisterCommand(
		"help",
		"shows a list of actual commands",
		b.help,
		manuals["help"],
	)

	b.RegisterCommand(
		"kill",
		"removes user from the chat",
		b.kill,
		manuals["kill"],
	)

	b.RegisterCommand(
		"respawn",
		"respawns the user",
		b.respawn,
		manuals["respawn"],
	)

	b.RegisterCommand(
		"rm",
		"removes message/messages",
		b.rm,
		manuals["rm"],
	)

	b.RegisterCommand(
		"wget",
		"displays info about the replied message",
		b.wget,
		manuals["wget"],
	)

	b.Tg.Init(b.onInit)
}

func (b *ZapBot) onInit(e general.User) {
	slog.Info(
		fmt.Sprintf("Bot %s started", e.FirstName),
	)
}

func (b *ZapBot) ProcessCommands(e general.Message) {
	if e.Text == "" {
		return
	}

	c, ok := core.StringToCommand(e.Text)

	if !ok {
		return
	}

	slog.Debug(
		fmt.Sprintf("Parsed command:\nCommand: %s\nargs: %v\netc: %s", c.Entry, c.Args, c.Etc),
	)

	command, ok := b.Commands[c.Entry]

	if !ok {
		return
	}

	command.Function(e)
}

func (b *ZapBot) RegisterCommand(entry string, description string, f func(general.Message), manual string) {
	c := Command{
		Description: description,
		Function:    f,
		Manual:      manual,
	}

	b.Commands[entry] = c
}

// You want to name your .md manual as the commands to make your life a little more simple
func (b *ZapBot) ReadManualsFrom(dir string) map[string]string {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)

	directory, err := os.ReadDir(path + dir)
	nameToMan := map[string]string{}

	if err != nil {
		slog.Error(err.Error())
		return nameToMan
	}

	for _, dirEntry := range directory {
		fileAndExt := strings.Split(dirEntry.Name(), ".")

		filename := fileAndExt[0]
		ext := fileAndExt[1]

		if ext != "md" {
			continue
		}

		manual, err := os.ReadFile(
			path + dir + "/" + dirEntry.Name(),
		)

		if err != nil {
			slog.Error(err.Error())
			return nameToMan
		}

		nameToMan[filename] = string(manual)
	}

	return nameToMan
}
