package core

import "slices"

type Chat struct {
	Users    map[int]*User
	Search   []*User
	Messages map[int]int
}

func (c *Chat) AddToSearch(id int) {
	c.Search = append(c.Search, c.Users[id])
}

func (c *Chat) RemoveFromSearch(id int) {
	// bro im tired... it's 4 am...
	newSearch := []*User{}

	for _, c := range c.Search {
		if c.Id != id {
			newSearch = append(newSearch, c)
		}
	}

	c.Search = newSearch
}

// returns companion id, equal interests
func (c *Chat) GetFirstCompanion(id int) (int, []int) {
	user := c.Users[id]

	for _, c := range c.Search {
		equalInterests := []int{}

		if len(user.Interests) != 0 || len(c.Interests) != 0 {
			for j := 0; j < len(c.Interests); j++ {
				if slices.Contains(user.Interests, c.Interests[j]) {
					equalInterests = append(equalInterests, c.Interests[j])
				}
			}

		}

		if c != user {
			if len(user.Interests) == 0 && len(c.Interests) == 0 {
				return c.Id, nil
			}

			if len(equalInterests) > 0 {
				return c.Id, equalInterests
			}
		}
	}

	return 0, nil
}

func (c *Chat) Connect(a int, b int) {
	c.Users[a].SetCompanion(b)
	c.Users[b].SetCompanion(a)

	c.RemoveFromSearch(a)
	c.RemoveFromSearch(b)
}

// disconnects a and a's companion (b) and returns b if exists
func (c *Chat) Disconnect(a int) int {
	b := c.Users[a].Companion

	c.Users[a].RemoveCompanion()

	for k := range c.Messages {
		if k == a || k == b {
			delete(c.Messages, k)
		}
	}

	if b != 0 {
		c.Users[b].RemoveCompanion()
		return b
	}

	return 0
}

func (c *Chat) Get(id int) int {
	return c.Users[id].Companion
}

func (c *Chat) AddMessage(a int, b int) {
	if c.Messages == nil {
		c.Messages = map[int]int{}
	}

	c.Messages[a] = b
}

func (c *Chat) GetMessageA(a int) int {
	v, ok := c.Messages[a]

	if ok {
		return v
	} else {
		return 0
	}
}

func (c *Chat) GetMessageB(b int) int {
	for k, v := range c.Messages {
		if v == b {
			return k
		}
	}

	return 0
}
