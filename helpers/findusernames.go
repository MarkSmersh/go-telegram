package helpers

import (
	"regexp"
)

func FindUsernames(s string) []string {
	r, _ := regexp.Compile(`@[a-zA-Z_]{5,}`)

	matches := r.FindAllStringSubmatch(s, -1)

	usernames := []string{}

	for _, v := range matches {
		usernames = append(usernames, v[0])
	}

	return usernames
}
