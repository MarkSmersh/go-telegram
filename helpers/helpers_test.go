package helpers

import "testing"

func TestFindUsernames(t *testing.T) {
	s := "Hi, I, @qekkk, wanted to say, that @izumm_risuet is a beautiful person."

	usernames := FindUsernames(s)

	if usernames[0] != "@qekkk" || usernames[1] != "@izumm_risuet" {
		t.Errorf("Wanted %s, %s, got %s, %s", "@qekkk", "@izumm_risuet", usernames[0], usernames[1])
	}
}
