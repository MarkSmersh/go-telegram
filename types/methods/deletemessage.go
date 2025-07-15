package methods

type DeleteMessage struct {
	ChatID    int `json:"chat_id"` // Use string to allow both int and @username in JSON form
	MessageID int `json:"message_id"`
}
