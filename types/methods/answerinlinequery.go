package methods

type AnswerInlineQuery struct {
	InlineQueryID string                    `json:"inline_query_id"`
	Results       []any                     `json:"results"`
	CacheTime     int                       `json:"cache_time,omitempty"`
	IsPersonal    bool                      `json:"is_personal,omitempty"`
	NextOffset    string                    `json:"next_offset,omitempty"`
	Button        *InlineQueryResultsButton `json:"button,omitempty"`
}

// You’ll need to define concrete result types depending on what you send.
// Telegram supports multiple kinds (article, gif, photo, etc.).
// This can be an interface or a union-like struct.

// Example button
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter string      `json:"start_parameter,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}
