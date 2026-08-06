package core

import (
	"github.com/MarkSmersh/go-telegram/types/methods"
)

type InlineBuilder struct {
	methods.AnswerInlineQuery
}

func NewInlineBuilder(id string) InlineBuilder {
	return InlineBuilder{
		AnswerInlineQuery: methods.AnswerInlineQuery{
			InlineQueryID: id,
		},
	}
}

func (b *InlineBuilder) addResult(obj any) {
	b.Results = append(b.Results, obj)
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

func (b *InlineBuilder) AddMpeg(mpeg InlineQueryResultMpeg4Gif) {
	b.addResult(mpeg)
}

type InlineQueryResultGif struct {
	Type              string `json:"type"` // must be "gif"
	ID                string `json:"id"`
	GifURL            string `json:"gif_url"`
	GifWidth          int    `json:"gif_width,omitempty"`
	GifHeight         int    `json:"gif_height,omitempty"`
	GifDuration       int    `json:"gif_duration,omitempty"`
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

func (b *InlineBuilder) AddGIF(gif InlineQueryResultGif) {
	b.addResult(gif)
}
