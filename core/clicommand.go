package core

import (
	"fmt"
	"log/slog"
	"regexp"
)

type CliCommand struct {
	Entry string
	Args  map[string]string
	Etc   string
}

func (c *CliCommand) GetArg(arg string) string {
	v, ok := c.Args[arg]
	if ok {
		return v
	} else {
		return ""
	}
}

func (c *CliCommand) IsArg(arg string) bool {
	for k := range c.Args {
		if k == arg {
			return true
		}
	}

	return false
}

func StringToCommand(s string) (CliCommand, bool) {
	c := CliCommand{}

	commandRegex, _ := regexp.Compile(`([a-z]{2,}) ?((?: ?(?:(?:--|-)[a-zA-Z-]+(?:[ |=][a-zA-Z0-9]+)?))*) ?([^\n]*)`)
	argsRegex, _ := regexp.Compile(`(?: ?(?:(?:--|-)([a-zA-Z-]+)(?:[ |=]([^ |-|(--)]+))?))`)

	argsFromCli := commandRegex.FindAllStringSubmatch(s, -1)

	slog.Debug(
		fmt.Sprintf("argsFromCli: %v", argsFromCli),
	)

	if len(argsFromCli) <= 0 {
		return c, false
	}

	command := argsFromCli[0][1]
	etc := argsFromCli[0][3]
	args := map[string]string{}

	slog.Debug(
		fmt.Sprintf("Command: %s; args: %v; etc: %s", command, args, etc),
	)

	argsFrom := argsRegex.FindAllStringSubmatch(argsFromCli[0][2], -1)

	for _, v := range argsFrom {
		args[v[1]] = v[2]
	}

	c.Entry = command
	c.Args = args
	c.Etc = etc

	slog.Debug(
		fmt.Sprintf("command: %v\n", c),
	)

	return c, true
}
