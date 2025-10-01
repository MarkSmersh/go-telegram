package zapbot

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	"github.com/MarkSmersh/go-telegram/components/rule34"
	"github.com/MarkSmersh/go-telegram/components/rule34/types/methods"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
)

func (b ZapBot) rule34(e core.Message) {
	rule34Token := os.Getenv("RULE34_API")
	rule34UserId := helpers.GetEnvInt("RULE34_USERID")

	if rule34Token == "" || rule34UserId == 0 {
		e.Reply("Very sorry, but module <code>rule34</code> is disabled due to an API token or a user id absence. Contact your administator and keep youself safe. You will traspass this.")
	}

	rule34 := rule34.NewRule34(rule34Token, rule34UserId)

	cli, _ := cli.NewCli(e.Raw().Text)

	count := 1
	tags := ""

	countOption, ok := cli.Get("c")

	if ok {
		count, ok = countOption.AtoiRange(1, 10)

		if !ok {
			fallbacks.OptionInvalidValue(e, countOption)
			return
		}
	}

	if e.ReplyToMessage != nil {
		tags = e.ReplyToMessage.Raw().Text
	}

	tags = cli.Etc

	posts, err := rule34.Posts(methods.Post{
		Limit: count,
		Tags:  tags,
	})

	if err != nil {
		slog.Error(err.Error())
		e.Reply(err.Error())
		return
	}

	reply := ""

	for _, p := range posts {
		reply += fmt.Sprintf("imgurl: %s\n", p.FileURL)
	}

	e.Reply(reply)
}
