package helpers

import (
	"regexp"
)

// Extracts telegram usernames (@example) from string into array ["example"]
func FindUsernames(s string) []string {
	r, _ := regexp.Compile(`@[a-zA-Z_]{5,}`)

	matches := r.FindAllStringSubmatch(s, -1)

	usernames := []string{}

	for _, v := range matches {
		usernames = append(usernames, v[0])
	}

	return usernames
}
