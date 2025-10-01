package methods

type DeleteMessages struct {
	ChatID int `json:"chat_id"` // Use string to allow both int and @username in JSON form
	// JSON Array
	MessageIDs string `json:"message_ids"`
}
