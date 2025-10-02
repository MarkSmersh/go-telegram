package methods

import "github.com/MarkSmersh/go-telegram/types/general"

type SendMediaGroup struct {
	ChatID                int    `json:"chat_id"`                            // Required. Target chat ID or username
	Media                 string `json:"media"`                              // Required. Array of media items (2–10)
	MessageThreadID       int64  `json:"message_thread_id,omitempty"`        // Optional. Thread ID (forum supergroups)
	DirectMessagesTopicID int64  `json:"direct_messages_topic_id,omitempty"` // Optional. Topic ID for direct messages chat

	BusinessConnectionID string                   `json:"business_connection_id,omitempty"` // Optional. Business connection ID
	DisableNotification  bool                     `json:"disable_notification,omitempty"`   // Optional. Send silently
	ProtectContent       bool                     `json:"protect_content,omitempty"`        // Optional. Protect from forwarding/saving
	AllowPaidBroadcast   bool                     `json:"allow_paid_broadcast,omitempty"`   // Optional. Paid broadcast bypass
	MessageEffectID      string                   `json:"message_effect_id,omitempty"`      // Optional. Message effect (private chats only)
	ReplyParameters      *general.ReplyParameters `json:"reply_parameters,omitempty"`       // Optional. Reply params
}
