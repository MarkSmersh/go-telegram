package methods

type UnbanChatMember struct {
	ChatID       int  `json:"chat_id"`                  // Unique ID or username of the target chat
	UserID       int  `json:"user_id"`                  // Unique ID of the target user
	OnlyIfBanned bool `json:"only_if_banned,omitempty"` // Do nothing if the user is not banned
}
