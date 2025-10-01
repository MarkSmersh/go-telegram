package core

import (
	"github.com/MarkSmersh/go-telegram/types/general"
)

type User struct {
	raw general.User
	tg  Telegram
}

func (t Telegram) NewUser(rawUser general.User) User {
	u := User{
		raw: rawUser,
		tg:  t,
	}

	return u
}

func (u User) Raw() general.User {
	return u.raw
}
