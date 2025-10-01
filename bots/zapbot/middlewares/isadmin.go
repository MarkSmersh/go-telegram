package middlewares

import "github.com/MarkSmersh/go-telegram/core"

func (m *Middleware) IsAdmin(e core.Message) bool {
	cm, _ := e.GetChatMember()

	if cm.IsCreator() || cm.IsAdmin() {
		return true
	} else {
		return false
	}
}
