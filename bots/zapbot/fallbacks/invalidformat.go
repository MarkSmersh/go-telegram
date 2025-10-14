package fallbacks

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
)

func InvalidFormat(e core.Message, c cli.Cli, comment string) {
	e.Reply(
		fmt.Sprintf("Invalid format of the command. %s\n\nRead the manual with <code>man %s</code>", comment, c.Entry),
	)
}
