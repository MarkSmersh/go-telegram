package core

import (
	"errors"
	"slices"
)

type UserModel struct {
	Id        int
	Interests []int
	Companion int
	Age       int
	Sex       int
}

func (u *UserModel) AddOrRemoveInterest(i int) {
	err := u.AddInterest(i)

	if err != nil {
		u.RemoveInterest(i)
	}
}

func (u *UserModel) AddInterest(i int) error {
	if u.Interests == nil {
		u.Interests = []int{}
	}

	if !slices.Contains(u.Interests, i) {
		u.Interests = append(u.Interests, i)
	} else {
		return errors.New("UserModel alredy has this interest")
	}

	return nil
}

func (u *UserModel) RemoveInterest(i int) {
	newInterests := []int{}

	for j := 0; j < len(u.Interests); j++ {
		if u.Interests[j] != i {
			newInterests = append(newInterests, u.Interests[j])
		}
	}

	u.Interests = newInterests
}

func (u *UserModel) SetSex(s int) {
	u.Sex = s
}

func (u *UserModel) SetAge(a int) {
	u.Sex = a
}

func (u *UserModel) SetCompanion(c int) {
	u.Companion = c
}

func (u *UserModel) RemoveCompanion() {
	u.Companion = 0
}
