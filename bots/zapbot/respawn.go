package zapbot

import (
	"fmt"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
)

func (b *ZapBot) respawn(e core.Message) {
	cli, _ := cli.NewCli(e.Raw().Text)

	target := 0

	if e.ReplyToMessage != nil {
		target = e.ReplyToMessage.From.Raw().ID
	}

	user, ok := cli.Get("u")

	if ok {
		id, err := b.ExtractUserId(user)

		if err != nil {
			fallbacks.OptionInvalidValue(e, user)
			return
		}

		target = id
	}

	victim, _ := e.Chat.GetMember(target)

	victim.SetChat(e.Chat)

	_, err := victim.Unban()

	if err != nil {
		e.Reply(fmt.Sprintf("User cannot be respawned.\n\n%s", err.Error()))
		return
	}

	e.Reply(fmt.Sprintf("Respawned %s. Invite him someone", victim.User.Raw().FirstName))
}
