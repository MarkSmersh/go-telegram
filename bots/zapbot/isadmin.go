package zapbot

import (
	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) isAdmin(c general.Chat, u general.User) bool {
	m, _ := b.Tg.GetChatMember(methods.GetChatMember{
		ChatID: c.ID,
		UserID: u.ID,
	})

	if m.Status == consts.CREATOR || m.Status == consts.ADMINISTRATOR {
		return true
	} else {
		return false
	}
}
