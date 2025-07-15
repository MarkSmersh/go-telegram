package anonchat

import (
	"fmt"
	"log/slog"

	"github.com/MarkSmersh/go-telegram/consts"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

var env, _ = helpers.GetEnv()

type AnonBot struct {
	Tg    core.Telegram
	Chat  core.Chat
	Users core.State[int, int]
}

func (b *AnonBot) Init() {
	b.Chat.Users = map[int]*core.User{}

	b.Tg.Eventer.Messages.Add(b.messageHandler)
	b.Tg.Eventer.Commands.Add(b.commandHandler)
	b.Tg.Eventer.CallbackQuery.Add(b.callbackHandler)
	b.Tg.Init(b.onInit)
}

func (b *AnonBot) onInit(me general.User) {
	slog.Info(
		fmt.Sprintf("Bot %s is started", me.FirstName),
	)
}

func (b *AnonBot) messageHandler(e general.Message) {
	if b.Users.Get(e.Chat.ID) == consts.StateConnected {
		req := methods.CopyMessage{
			ChatID:     b.Chat.Get(e.Chat.ID),
			FromChatID: e.Chat.ID,
			MessageID:  e.MessageID,
		}

		if e.ReplyToMessage != nil {

			messageId := b.Chat.GetMessageA(e.ReplyToMessage.MessageID)
			from := b.Chat.Get(e.Chat.ID)

			if e.ReplyToMessage.From.ID == e.From.ID {
				messageId = b.Chat.GetMessageB(e.ReplyToMessage.MessageID)
			}

			rp := general.ReplyParameters{
				MessageID:                messageId,
				ChatID:                   from,
				AllowSendingWithoutReply: true,
			}

			req.ReplyParameters = rp.ToJSON()
		}

		mes, _ := b.Tg.CopyMessage(req)

		b.Chat.AddMessage(mes.MessageID, e.MessageID)
	}
}

func (b *AnonBot) commandHandler(e general.Message) {
	_, ok := b.Chat.Users[e.Chat.ID]

	if !ok {
		b.Chat.Users[e.Chat.ID] = &core.User{Id: e.Chat.ID, Interests: []int{}, Age: 0, Companion: 0, Sex: 0}
	}

	b.start(e)

	b.chatSearch(e)

	b.stop(e)

	b.ping(e)

	b.interests(e)
}

func (b *AnonBot) callbackHandler(e general.CallbackQuery) {
	_, ok := b.Chat.Users[e.From.ID]

	if !ok {
		b.Chat.Users[e.From.ID] = &core.User{Id: e.From.ID, Interests: []int{}, Age: 0, Companion: 0, Sex: 0}
	}

	b.userInterests(e)
}
