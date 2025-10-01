package core

import (
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

type Chat struct {
	raw general.Chat
	tg  Telegram
}

func (t Telegram) NewChat(rawChat general.Chat) Chat {
	c := Chat{
		raw: rawChat,
		tg:  t,
	}

	return c
}

func (c Chat) GetMember(userId int) (ChatMember, error) {
	chatMember, err := c.tg.GetChatMember(methods.GetChatMember{
		ChatID: c.Raw().ID,
		UserID: userId,
	})

	return chatMember, err
}

func (c Chat) Raw() general.Chat {
	return c.raw
}
