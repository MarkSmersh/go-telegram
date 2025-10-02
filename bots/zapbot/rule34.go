package zapbot

import (
	"log/slog"
	"math"
	"math/rand"
	"os"

	"github.com/MarkSmersh/go-telegram/bots/zapbot/fallbacks"
	"github.com/MarkSmersh/go-telegram/components/cli"
	rule34component "github.com/MarkSmersh/go-telegram/components/rule34"
	rule34general "github.com/MarkSmersh/go-telegram/components/rule34/types/general"
	rule34methods "github.com/MarkSmersh/go-telegram/components/rule34/types/methods"
	"github.com/MarkSmersh/go-telegram/core"
	"github.com/MarkSmersh/go-telegram/helpers"
	"github.com/MarkSmersh/go-telegram/types/general"
)

func (b ZapBot) rule34(e core.Message) {
	rule34Token := os.Getenv("RULE34_API")
	rule34UserId := helpers.GetEnvInt("RULE34_USERID")

	if rule34Token == "" || rule34UserId == 0 {
		e.Reply("Very sorry, but module <code>rule34</code> is disabled due to an API token or a user id absence. Contact your administator and keep youself safe. You will traspass this.")
	}

	rule34 := rule34component.NewRule34(rule34Token, rule34UserId)

	cli, _ := cli.NewCli(e.Raw().Text)

	count := 1
	tags := ""
	page := 0

	countOption, ok := cli.Get("c")

	if ok {
		count, ok = countOption.AtoiRange(2, 10)

		if !ok {
			fallbacks.OptionInvalidValue(e, countOption)
			return
		}
	}

	pageOption, ok := cli.Get("p")

	if ok {
		page, ok = pageOption.AtoiRange(1, 1000000)

		if !ok {
			fallbacks.OptionInvalidValue(e, pageOption)
			return
		}
	}

	if e.ReplyToMessage != nil {
		tags = e.ReplyToMessage.Raw().Text
	}

	tags = cli.Etc

	posts := []rule34general.Post{}

	if cli.Exists("r") {
		tempPosts, err := rule34.Posts(rule34methods.Post{
			Limit: 1000,
			Tags:  tags,
			PID:   page,
		})

		if err != nil {
			e.Reply(err.Error())
			return
		}

		if len(tempPosts) <= count {
			posts = tempPosts
		} else {
			rand.Shuffle(len(tempPosts), func(i, j int) {
				tempPosts[i], tempPosts[j] = tempPosts[j], tempPosts[i]
			})

			posts = tempPosts[:count]
		}

	} else {
		var err error

		posts, err = rule34.Posts(rule34methods.Post{
			Limit: count,
			Tags:  tags,
			PID:   page,
		})

		if err != nil {
			slog.Error(err.Error())
			e.Reply(err.Error())
			return
		}
	}

	if len(posts) <= 0 {
		e.Reply("Nothing found by your prompt. Out of bounds.")
		return
	}

	media := core.NewInputMediaBuilder()

	for _, p := range posts {
		postTagsLen := int(
			math.Min(1000, float64(len(p.Tags))),
		)

		media.AddPhoto(general.InputMediaPhoto{
			Type:       "photo",
			Media:      p.FileURL,
			Caption:    p.Tags[:postTagsLen],
			HasSpoiler: true,
		})
	}

	e.SendMediaGroup(media)
}
