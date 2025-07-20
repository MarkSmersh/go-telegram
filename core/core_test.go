package core

import (
	"testing"
)

func TestStringToCommand(t *testing.T) {
	c, _ := StringToCommand("command")

	if c.Entry != "command" {
		t.Errorf(`StringToCommand("command"), want command, got %s`, c.Etc)
	}

	c, _ = StringToCommand("command -f")

	if !c.IsArg("f") {
		t.Errorf(`StringToCommand("command -f"), want true, got false`)
	}

	c, _ = StringToCommand("command -f arg1")

	if arg := c.GetArg("f"); arg != "arg1" {
		t.Errorf(`StringToCommand("command -f arg1"), want arg1, got %s`, arg)
	}

	c, _ = StringToCommand("command -f arg1 -fd arg2 -fd=arg3 etc")

	if etc := c.Etc; etc != "etc" {
		t.Errorf(`StringToCommand("command -f arg1 etc"), want etc, got %s`, etc)
	}
}
