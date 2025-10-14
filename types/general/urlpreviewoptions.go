package general

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`        // Optional. True if the link preview is disabled
	URL              string `json:"url,omitempty"`                // Optional. URL to use for the link preview
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"` // Optional. True if media should be shrunk
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"` // Optional. True if media should be enlarged
	ShowAboveText    bool   `json:"show_above_text,omitempty"`    // Optional. True if the link preview should appear above the message text
}
