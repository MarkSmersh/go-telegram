package zapbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/MarkSmersh/go-telegram/components/cli"
)

func (b ZapBot) RegisterUsername(id int, username string) {
	b.Rdb.Set(
		context.Background(),
		fmt.Sprintf("userid:%d", id),
		username,
		0,
	)
	b.Rdb.Set(
		context.Background(),
		fmt.Sprintf("username:%s", username),
		id,
		0,
	)
}

func (b ZapBot) GetUsername(id int) string {
	res, err := b.Rdb.Get(
		context.Background(),
		fmt.Sprintf("userid:%d", id),
	).Result()

	if err != nil {
		return "nil"
	} else {
		return res
	}
}

func (b ZapBot) GetUserMention(id int) string {
	username, err := b.Rdb.Get(
		context.Background(),
		fmt.Sprintf("userid:%d", id),
	).Result()

	if err == nil {
		return fmt.Sprintf("[@%s](tg://user?id=%d)", username, id)
	} else {
		return fmt.Sprintf("[id%s](tg://user?id=%d)", "smb", id)
	}
}

func (b ZapBot) ExtractUserId(opt cli.Option) (int, error) {
	slog.Debug(
		fmt.Sprintf("opt: %v", opt),
	)

	if opt.Value[0] == '@' {
		username := strings.Split(opt.Value, "@")[1]
		res, err := b.Rdb.Get(context.Background(), "username:"+username).Result()

		if err != nil {
			return 0, errors.New("There is no user with such username.")
		}

		id, _ := strconv.Atoi(res)

		return id, nil
	}

	id, err := strconv.Atoi(opt.Value)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Invalid argument for an option %s", opt.Prefix))
	}

	// if id < 1 {
	// 	return 0, errors.New(fmt.Sprintf("Invalid argument for an option %s", opt.Prefix))
	// }

	return id, nil
}
