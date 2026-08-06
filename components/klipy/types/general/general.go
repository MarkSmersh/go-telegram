package general

import "encoding/json"

type Response struct {
	Result bool     `json:"result"`
	Data   PageData `json:"data"`
}

type PageData struct {
	Data        json.RawMessage `json:"data"`
	CurrentPage int             `json:"current_page"`
	PerPage     int             `json:"per_page"`
	HasNext     bool            `json:"has_next"`
}

type Item struct {
	ID          int64    `json:"id"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	File        File     `json:"file"`
	Tags        []string `json:"tags"`
	Type        string   `json:"type"`
	BlurPreview string   `json:"blur_preview"`
}

type File struct {
	HD SizeGroup `json:"hd"`
	MD SizeGroup `json:"md"`
	SM SizeGroup `json:"sm"`
	XS SizeGroup `json:"xs"`
}

type SizeGroup struct {
	GIF  Media `json:"gif"`
	WebP Media `json:"webp"`
	JPG  Media `json:"jpg"`
	MP4  Media `json:"mp4"`
	WebM Media `json:"webm"`
}

type Media struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"size"`
}
