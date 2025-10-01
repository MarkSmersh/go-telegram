package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
)

func (b *ZapBot) man(e core.Message) {
	cli, _ := cli.NewCli(e.Raw().Text)

	if cli.Etc == "" {
		e.Reply("What manual do you want? For example, try <code>man man</code> ")
		return
	}

	c, ok := b.CommandsMng.Get(cli.Etc)

	if !ok {
		e.Reply(fmt.Sprintf("No manual entry for <code>%s</code>", cli.Etc))
		return
	}

	e.Reply(fmt.Sprintf("<code>%s</code> - <i>%s</i>\n\n%s", cli.Etc, c.Description, c.Manual))
}
