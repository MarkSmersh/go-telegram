package core

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

type Message struct {
	raw                general.Message
	tg                 Telegram
	parseMode          string
	disableLinkPreview bool

	From *User `json:"from,omitempty"`
	// SenderChat        *Chat    `json:"sender_chat,omitempty"`
	// SenderBusinessBot *User    `json:"sender_business_bot,omitempty"`
	Chat           Chat     `json:"chat"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	ViaBot         *User    `json:"via_bot,omitempty"`
	// NewChatMembers    []User   `json:"new_chat_members,omitempty"`
	// LeftChatMember    *User    `json:"left_chat_member,omitempty"`
}

func (t Telegram) NewMessage(raw general.Message) Message {
	// TODO: Add these and others models

	msg := Message{
		raw:                raw,
		tg:                 t,
		parseMode:          "",
		disableLinkPreview: false,

		// From:           &from,
		// Chat:           chat,
		// ReplyToMessage: &replytomessage,
		// ViaBot:         &viabot,
	}

	if raw.From != nil {
		from := t.NewUser(*raw.From)
		msg.From = &from
	}

	msg.Chat = t.NewChat(raw.Chat)

	if raw.ReplyToMessage != nil {
		replytomessage := t.NewMessage(*raw.ReplyToMessage)
		msg.ReplyToMessage = &replytomessage
	}

	if raw.ViaBot != nil {
		viabot := t.NewUser(*raw.ViaBot)
		msg.ViaBot = &viabot
	}

	return msg
}

func (m Message) Reply(text string) (Message, error) {
	msg, err := m.tg.SendMessage(methods.SendMessage{
		Text:      text,
		ParseMode: m.parseMode,
		ChatID:    m.Chat.Raw().ID,
		LinkPreviewOptions: &general.LinkPreviewOptions{
			IsDisabled: m.disableLinkPreview,
		},
	})

	return msg, err
}

func (m *Message) SetParseMode(parseMode string) {
	availableParseModes := []string{
		consts.Markdown,
		consts.MarkdownV2,
		consts.HTML,
	}

	if !slices.Contains(availableParseModes, parseMode) {
		slog.Warn(
			fmt.Sprintf(
				"ParseMode %s doesn't exist. Available options are: %s",
				parseMode,
				strings.Join(availableParseModes, ""),
			),
		)
	}

	m.parseMode = parseMode
}

func (m *Message) ClearParseMode() {
	m.parseMode = ""
}

func (m Message) DeleteMessage() (bool, error) {
	return m.tg.DeleteMessage(methods.DeleteMessage{
		ChatID:    m.raw.Chat.ID,
		MessageID: m.raw.MessageID,
	})
}

func (m Message) GetChatMember() (ChatMember, error) {
	chatMember, err := m.tg.GetChatMember(methods.GetChatMember{
		ChatID: m.Chat.Raw().ID,
		UserID: m.From.Raw().ID,
	})

	chatMember.SetChat(m.Chat)
	return chatMember, err
}

func (m Message) SendMediaGroup(mediaBuilder InputMediaBuilder) ([]Message, error) {
	return m.tg.SendMediaGroup(methods.SendMediaGroup{
		ChatID: m.Chat.Raw().ID,
		Media:  mediaBuilder.ToJSON(),
	})
}

func (m *Message) DisableLinkPreview() {
	m.disableLinkPreview = true
}

func (m *Message) EnableLinkPreview() {
	m.disableLinkPreview = false
}

func (m Message) Raw() general.Message {
	return m.raw
}
