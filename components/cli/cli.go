package cli

import (
	"fmt"
	"log/slog"
	"regexp"
)

type Cli struct {
	Entry string
	Args  []Option
	Etc   string
}

func NewCli(s string) (Cli, bool) {
	c := Cli{}

	commandRegex, _ := regexp.Compile(`([a-z0-9]{2,}) ?((?: ?(?:(?:--|-)[a-zA-Z-]+(?:[ |=][a-zA-Z0-9@-_]+)?))*) ?([^\n]*)`)
	argsRegex, _ := regexp.Compile(`(?: ?(?:(?:--|-)([a-zA-Z]+)(?:[ |=]([^ -]+))?))`)

	argsFromCli := commandRegex.FindAllStringSubmatch(s, -1)

	slog.Debug(
		fmt.Sprintf("argsFromCli: %v", argsFromCli),
	)

	if len(argsFromCli) <= 0 {
		return c, false
	}

	command := argsFromCli[0][1]
	etc := argsFromCli[0][3]
	// args := map[string]string{}
	args := []Option{}

	slog.Debug(
		fmt.Sprintf("Command: %s; args: %v; etc: %s", command, args, etc),
	)

	argsFrom := argsRegex.FindAllStringSubmatch(argsFromCli[0][2], -1)

	for _, v := range argsFrom {
		a := Option{
			Prefix: v[1],
			Value:  v[2],
		}

		args = append(args, a)
	}

	c.Entry = command
	c.Args = args
	c.Etc = etc

	slog.Debug(
		fmt.Sprintf("command: %v\n", c),
	)

	return c, true
}

func (c *Cli) Get(opt string) (Option, bool) {
	var v *Option = nil

	for _, a := range c.Args {
		if opt == a.Prefix {
			v = &a
		}
	}

	if v != nil {
		return *v, true
	} else {
		return Option{}, false
	}
}

func (c *Cli) Exists(opt string) bool {
	_, ok := c.Get(opt)
	return ok
}
