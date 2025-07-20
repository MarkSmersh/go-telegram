package general

type ChatMember struct {
	Status string `json:"status"` // creator, administrator, member, restricted, left, kicked
	User   User   `json:"user"`   // shared by all

	// Common optional fields
	CustomTitle string `json:"custom_title,omitempty"` // owner/admin
	UntilDate   int64  `json:"until_date,omitempty"`   // member, restricted, banned

	// Owner/Admin-specific
	IsAnonymous         bool `json:"is_anonymous,omitempty"`           // owner/admin
	CanBeEdited         bool `json:"can_be_edited,omitempty"`          // admin
	CanManageChat       bool `json:"can_manage_chat,omitempty"`        // admin
	CanDeleteMessages   bool `json:"can_delete_messages,omitempty"`    // admin
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"` // admin
	CanRestrictMembers  bool `json:"can_restrict_members,omitempty"`   // admin
	CanPromoteMembers   bool `json:"can_promote_members,omitempty"`    // admin
	CanChangeInfo       bool `json:"can_change_info,omitempty"`        // admin/restricted
	CanInviteUsers      bool `json:"can_invite_users,omitempty"`       // admin/restricted
	CanPostStories      bool `json:"can_post_stories,omitempty"`       // admin
	CanEditStories      bool `json:"can_edit_stories,omitempty"`       // admin
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`     // admin

	// Optional fields (channels/supergroups only)
	CanPostMessages bool `json:"can_post_messages,omitempty"` // admin (optional)
	CanEditMessages bool `json:"can_edit_messages,omitempty"` // admin (optional)
	CanPinMessages  bool `json:"can_pin_messages,omitempty"`  // admin/restricted (optional)
	CanManageTopics bool `json:"can_manage_topics,omitempty"` // admin/restricted (optional)

	// Restricted-specific
	IsMember              bool `json:"is_member,omitempty"` // restricted only
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
}
