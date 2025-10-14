package fallbacks

import (
	"github.com/MarkSmersh/go-telegram/core"
)

func UnableToProceed(e core.Message) {
	e.Reply("Unable to process your request.")
}
