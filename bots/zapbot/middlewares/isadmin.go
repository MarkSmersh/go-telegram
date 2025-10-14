package middlewares

import "github.com/MarkSmersh/go-telegram/core"

func (m *Middleware) IsAdmin(e core.Message) bool {
	if e.From.Raw().ID == e.Chat.Raw().ID {
		return true
	}

	cm, _ := e.GetChatMember()

	if cm.IsCreator() || cm.IsAdmin() {
		return true
	} else {
		return false
	}
}
