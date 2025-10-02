package core

import (
	"encoding/json"

	"github.com/MarkSmersh/go-telegram/types/general"
)

type InputMediaBuilder struct {
	photos []general.InputMediaPhoto
}

func NewInputMediaBuilder() InputMediaBuilder {
	return InputMediaBuilder{
		photos: []general.InputMediaPhoto{},
	}
}

func (i *InputMediaBuilder) AddPhoto(photo general.InputMediaPhoto) {
	i.photos = append(i.photos, photo)
}

func (i InputMediaBuilder) ToJSON() string {
	inputs := []any{}

	for _, p := range i.photos {
		inputs = append(inputs, p)
	}

	bytes, _ := json.Marshal(inputs)
	return string(bytes)
}
