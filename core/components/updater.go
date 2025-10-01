package components

import (
	"github.com/MarkSmersh/go-telegram/types/general"
)

type Updater struct {
	Messages      Caller[general.Message]
	Commands      Caller[general.Message]
	InlineQuery   Caller[general.InlineQuery]
	CallbackQuery Caller[general.CallbackQuery]
}
