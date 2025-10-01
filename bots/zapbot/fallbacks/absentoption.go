package fallbacks

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
)

func OptionIsAbsent(e core.Message, cliOpt cli.Option) {
	e.Reply(fmt.Sprintf("Argument <code>%s</code> is absent.", cliOpt.Prefix))
}
