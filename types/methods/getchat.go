package methods

type GetChat struct {
	ChatID any `json:"chat_id"` // Required: int64 or string (@channelusername)
}
