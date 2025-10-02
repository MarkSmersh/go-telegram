package general

type InputMediaPhoto struct {
	Type      string `json:"type"`                 // Must be "photo"
	Media     string `json:"media"`                // File ID, URL, or attach://<file_attach_name>
	Caption   string `json:"caption,omitempty"`    // Optional. Caption text (0-1024 chars)
	ParseMode string `json:"parse_mode,omitempty"` // Optional. Parse mode (Markdown or HTML)
	// CaptionEntities      []MessageEntity `json:"caption_entities,omitempty"`          // Optional. Entities in caption (instead of parse_mode)
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"` // Optional. Show caption above the media
	HasSpoiler            bool `json:"has_spoiler,omitempty"`              // Optional. Spoiler animation
}
