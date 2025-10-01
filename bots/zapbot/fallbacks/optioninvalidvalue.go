package fallbacks

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
)

func OptionInvalidValue(e core.Message, cliOpt cli.Option) {
	e.Reply(fmt.Sprintf("value for an option <code>%s</code> is invalid.", cliOpt.Prefix))
}
