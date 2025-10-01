package methods

// RequestParams holds query parameters for the API
type Post struct {
	Limit int    `json:"limit,omitempty"` // max 1000
	PID   int    `json:"pid,omitempty"`   // page number
	Tags  string `json:"tags,omitempty"`  // space-separated tags
	CID   int64  `json:"cid,omitempty"`   // change ID (unix time)
	ID    int    `json:"id,omitempty"`    // post id
}
