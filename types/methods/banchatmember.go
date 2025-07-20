package methods

type BanChatMember struct {
	ChatID         int  `json:"chat_id"`                   // int64 or string (@channelusername)
	UserID         int  `json:"user_id"`                   // user to ban
	UntilDate      int  `json:"until_date,omitempty"`      // optional unix time
	RevokeMessages bool `json:"revoke_messages,omitempty"` // optional
}
