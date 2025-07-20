package general

type ChatFullInfo struct {
	ID               int64  `json:"id"`   // Unique identifier for this chat
	Type             string `json:"type"` // Chat type: “private”, “group”, “supergroup” or “channel”
	Title            string `json:"title,omitempty"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	IsForum          bool   `json:"is_forum,omitempty"`
	AccentColorID    int    `json:"accent_color_id"`
	MaxReactionCount int    `json:"max_reaction_count"`
	// Photo                                *ChatPhoto              `json:"photo,omitempty"`
	ActiveUsernames []string `json:"active_usernames,omitempty"`
	// Birthdate                            *Birthdate              `json:"birthdate,omitempty"`
	// BusinessIntro                        *BusinessIntro          `json:"business_intro,omitempty"`
	// BusinessLocation                     *BusinessLocation       `json:"business_location,omitempty"`
	// BusinessOpeningHours                 *BusinessOpeningHours   `json:"business_opening_hours,omitempty"`
	PersonalChat *Chat `json:"personal_chat,omitempty"`
	// AvailableReactions                   []ReactionType          `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            string   `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               int      `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string   `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           string   `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64    `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string   `json:"bio,omitempty"`
	HasPrivateForwards                 bool     `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool     `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool     `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool     `json:"join_by_request,omitempty"`
	Description                        string   `json:"description,omitempty"`
	InviteLink                         string   `json:"invite_link,omitempty"`
	PinnedMessage                      *Message `json:"pinned_message,omitempty"`
	// Permissions                          *ChatPermissions        `json:"permissions,omitempty"`
	// AcceptedGiftTypes                    *AcceptedGiftTypes      `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia             bool   `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                int    `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount         int    `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime        int    `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled bool   `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers             bool   `json:"has_hidden_members,omitempty"`
	HasProtectedContent          bool   `json:"has_protected_content,omitempty"`
	HasVisibleHistory            bool   `json:"has_visible_history,omitempty"`
	StickerSetName               string `json:"sticker_set_name,omitempty"`
	CanSetStickerSet             bool   `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName    string `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                 int64  `json:"linked_chat_id,omitempty"`
	// Location                             *ChatLocation           `json:"location,omitempty"`
}
