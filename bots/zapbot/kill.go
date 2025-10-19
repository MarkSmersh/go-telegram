package zapbot

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

func (b *ZapBot) kill(e core.Message) {
	cli, _ := cli.NewCli(e.Raw().Text)

	target := 0
	permanent := false

	if cli.Exists("a") {
		e.Reply("GROUP WILL BE WIPED IN 30 SECONDS.")

		for i := 30; i > 0; i-- {
			e.Reply(
				fmt.Sprintf("%d...", i),
			)
			time.Sleep(1 * time.Second)
		}

		e.Reply("...THY END IS NOW")

		return
	}

	if e.ReplyToMessage != nil {
		target = e.ReplyToMessage.Raw().From.ID
	}

	if cli.Exists("f") || cli.Exists("ban") {
		permanent = true
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

	if target <= 0 {
		e.Reply("Option User ID is missing. Reply to message's owner you want to be banned. Also you can use -u or --user-id to set user'd is manually.")

		return
	}

	victim, _ := b.Tg.GetChatMember(methods.GetChatMember{
		ChatID: e.Chat.Raw().ID,
		UserID: target,
	})

	victim.SetChat(e.Chat)

	if victim.IsCreator() {
		e.Reply("How about KYS?")
		return
	}

	if victim.IsAdmin() {
		e.Reply("You cannot kill admin.")
		return
	}

	var err error

	if permanent {
		_, err = victim.PermBan()
		e.Reply(fmt.Sprintf("- Well... done... %s...\n- Farewell, %s.", e.From.Raw().FirstName, victim.Raw().User.FirstName))
	} else {
		_, err = victim.Kick()
		e.Reply(fmt.Sprintf("Killed %s.", victim.Raw().User.FirstName))
	}

	if err != nil {
		e.Reply(fmt.Sprintf("User cannot be removed.\n\n%s", err.Error()))
		return
	}
}
