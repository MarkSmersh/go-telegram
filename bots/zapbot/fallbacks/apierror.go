package fallbacks

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/core"
)

func ApiError(e core.Message, error string, comment string) {
	e.Reply(fmt.Sprintf("%s.\n\n%s", comment, error))
}
