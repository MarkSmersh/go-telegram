package command

import (
	"log/slog"
	"os"
	// "path/filepath"
	// "runtime"
	"strings"

	"github.com/MarkSmersh/go-telegram/core"
)

type CommandsManager struct {
	commands map[string]Command
}

func NewCommandsManager() CommandsManager {
	return CommandsManager{
		commands: map[string]Command{},
	}
}

func (c *CommandsManager) RegisterCommand(entry string, description string, f func(core.Message), manual string, middlewares ...func(core.Message) bool) {
	command := Command{
		Description: description,
		Function:    f,
		Middlewares: middlewares,
		Manual:      manual,
	}

	c.commands[entry] = command
}

func (c *CommandsManager) Commands() map[string]Command {
	return c.commands
}

func (c *CommandsManager) Get(prefix string) (Command, bool) {
	cmd, ok := c.commands[prefix]
	return cmd, ok
}

// You want to name your .md manual as the commands to make your life a little more simple
func ReadManualsFrom(dir string) map[string]string {
	directory, err := os.ReadDir(dir)
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
			/* path + */ dir + "/" + dirEntry.Name(),
		)

		if err != nil {
			slog.Error(err.Error())
			return nameToMan
		}

		nameToMan[filename] = string(manual)
	}

	return nameToMan
}
