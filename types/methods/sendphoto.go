package methods

import "github.com/MarkSmersh/go-telegram/types/general"

type SendPhoto struct {
	BusinessConnectionID  string `json:"business_connection_id,omitempty"`   // Optional. Business connection ID
	ChatID                int    `json:"chat_id"`                            // Required. Target chat ID or username (@channelusername)
	MessageThreadID       int64  `json:"message_thread_id,omitempty"`        // Optional. Forum topic thread ID
	DirectMessagesTopicID int64  `json:"direct_messages_topic_id,omitempty"` // Optional. Direct messages topic ID
	Photo                 string `json:"photo"`                              // Required. Photo to send (file_id, URL, or InputFile)
	Caption               string `json:"caption,omitempty"`                  // Optional. Caption text
	ParseMode             string `json:"parse_mode,omitempty"`               // Optional. Parse mode for entities
	// CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities in caption
	ShowCaptionAboveMedia bool   `json:"show_caption_above_media,omitempty"` // Optional. Show caption above media
	HasSpoiler            bool   `json:"has_spoiler,omitempty"`              // Optional. Mark photo as spoiler
	DisableNotification   bool   `json:"disable_notification,omitempty"`     // Optional. Silent message
	ProtectContent        bool   `json:"protect_content,omitempty"`          // Optional. Protect content from forwarding
	AllowPaidBroadcast    bool   `json:"allow_paid_broadcast,omitempty"`     // Optional. Paid broadcast bypass
	MessageEffectID       string `json:"message_effect_id,omitempty"`        // Optional. Message effect ID
	// SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"` // Optional. Suggested post parameters (for DMs)
	ReplyParameters *general.ReplyParameters `json:"reply_parameters,omitempty"` // Optional. Reply parameters
	ReplyMarkup     string                   `json:"reply_markup,omitempty"`     // Optional. InlineKeyboard, ReplyKeyboard, Remove, or ForceReply
}
