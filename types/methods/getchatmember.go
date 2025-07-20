package methods

type GetChatMember struct {
	ChatID int `json:"chat_id"` // or string, depending on use case
	UserID int `json:"user_id"`
}
