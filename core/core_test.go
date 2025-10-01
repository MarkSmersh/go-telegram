package core

import (
	"testing"

	"github.com/MarkSmersh/go-telegram/components/cli"
)

func TestNewCli(t *testing.T) {
	c, _ := cli.NewCli("command")

	if c.Entry != "command" {
		t.Errorf(`NewCli("command"), want command, got %s`, c.Etc)
	}

	c, _ = cli.NewCli("command -f")

	if _, ok := c.Get("f"); !ok {
		t.Errorf(`NewCli("command -f"), want true, got false`)
	}

	c, _ = cli.NewCli("command -f arg1")

	if arg, _ := c.Get("f"); arg.Value != "arg1" {
		t.Errorf(`NewCli("command -f arg1"), want arg1, got %s`, arg)
	}

	c, _ = cli.NewCli("command -f arg1 -fd arg2 -fd=arg3 etc")

	if etc := c.Etc; etc != "etc" {
		t.Errorf(`NewCli("command -f arg1 etc"), want etc, got %s`, etc)
	}
}
