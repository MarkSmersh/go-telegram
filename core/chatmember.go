package core

import (
	"errors"
	"log/slog"
	"time"

	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

type ChatMember struct {
	raw general.ChatMember
	tg  Telegram

	User User
	Chat *Chat

	RevokeMessages bool
}

func (t Telegram) NewChatMember(rawChatMember general.ChatMember, chat *Chat) ChatMember {
	return ChatMember{
		raw:  rawChatMember,
		tg:   t,
		User: t.NewUser(rawChatMember.User),
		Chat: chat,
	}
}

// if unixUntil equals zero, ban will be permanent
func (m ChatMember) Ban(unixUntil int) (bool, error) {
	if m.Chat == nil {
		err := errors.New("Chat isn't set for this ChatMember. To do it use method SetChat.")
		slog.Error(err.Error())
		return false, err
	}

	return m.tg.BanChatMember(methods.BanChatMember{
		ChatID:    m.Chat.Raw().ID,
		UserID:    m.User.Raw().ID,
		UntilDate: unixUntil,
	})
}

func (m ChatMember) PermBan() (bool, error) {
	return m.Ban(0)
}

// bans for 60s. This is the only seem way to "kick"
func (m ChatMember) Kick() (bool, error) {
	untilDate := time.Now().UnixMilli() + 60
	return m.Ban(int(untilDate))
}

func (m ChatMember) Unban() (bool, error) {
	return m.tg.UnbanChatMember(methods.UnbanChatMember{
		ChatID: m.Chat.Raw().ID,
		UserID: m.User.Raw().ID,
	})
}

func (m *ChatMember) SetChat(chat Chat) {
	m.Chat = &chat
}

// creator, administrator, member, restricted, left, kicked
func (m ChatMember) HasStatus(status string) bool {
	return m.raw.Status == status
}

func (m ChatMember) IsCreator() bool {
	return m.HasStatus("creator")
}

func (m ChatMember) IsAdmin() bool {
	return m.HasStatus("administrator")
}

func (m ChatMember) IsMember() bool {
	return m.HasStatus("member")
}

func (m ChatMember) IsRestricted() bool {
	return m.HasStatus("restricted")
}

func (m ChatMember) IsLeft() bool {
	return m.HasStatus("left")
}

func (m ChatMember) IsKicked() bool {
	return m.HasStatus("kicked")
}

func (m ChatMember) Raw() general.ChatMember {
	return m.raw
}
