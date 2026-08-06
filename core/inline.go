package core

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/MarkSmersh/go-telegram/types/methods"
)

type InlineBuilder struct {
	methods.AnswerInlineQuery
	// gifs []InlineQueryResultMpeg4Gif
	objects []any
}

func NewInlineBuilder(id string) InlineBuilder {
	return InlineBuilder{
		AnswerInlineQuery: methods.AnswerInlineQuery{
			InlineQueryID: id,
		},
	}
}

type InlineQueryResultMpeg4Gif struct {
	Type              string `json:"type"` // must be "mpeg4_gif"
	ID                string `json:"id"`
	Mpeg4URL          string `json:"mpeg4_url"`
	Mpeg4Width        int    `json:"mpeg4_width,omitempty"`
	Mpeg4Height       int    `json:"mpeg4_height,omitempty"`
	Mpeg4Duration     int    `json:"mpeg4_duration,omitempty"`
	ThumbnailURL      string `json:"thumbnail_url"`
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
	Title             string `json:"title,omitempty"`
	Caption           string `json:"caption,omitempty"`
	ParseMode         string `json:"parse_mode,omitempty"`
	// CaptionEntities        []MessageEntity        `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// ReplyMarkup            *InlineKeyboardMarkup  `json:"reply_markup,omitempty"`
	// InputMessageContent    InputMessageContent    `json:"input_message_content,omitempty"`
}

func (b *InlineBuilder) addObject(obj any) {
	b.objects = append(b.objects, obj)

	encoded, err := json.Marshal(b.objects)

	if err != nil {
		slog.Error(err.Error())
	}

	b.Results = string(encoded)

	slog.Debug(fmt.Sprintf("END: %d", len(b.objects)))
}

func (b *InlineBuilder) AddGIF(gif InlineQueryResultMpeg4Gif) {
	b.addObject(gif)
}
