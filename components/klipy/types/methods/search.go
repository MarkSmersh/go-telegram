package methods

type Search struct {
	Page          int    `json:"page,omitempty"`
	PerPage       int    `json:"per_page,omitempty"`
	Q             string `json:"q,omitempty"`
	CustomerID    string `json:"customer_id,omitempty"`
	Locale        string `json:"locale,omitempty"`
	ContentFilter string `json:"content_filter,omitempty"`
}
