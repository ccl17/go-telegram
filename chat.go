package telegram

import (
	"context"
)

type Chat struct {
	ID                  int      `json:"id"`
	Type                string   `json:"type"`
	Title               string   `json:"title"`
	UserName            string   `json:"username"`
	FirstName           string   `json:"first_name"`
	LastName            string   `json:"last_name"`
	AllMembersAreAdmins bool     `json:"all_members_are_administrators"`
	Description         string   `json:"description,omitempty"`
	InviteLink          string   `json:"invite_link,omitempty"`
	PinnedMessage       *Message `json:"pinned_message"`
	// Photo               *ChatPhoto `json:"photo"`
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`
	UntilDate             int    `json:"until_date,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

type ChatOptions struct {
	ChatId int `json:"chat_id"`
}

func (c *BotClient) SendChatAction(ctx context.Context, options SendChatActionOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSendChatAction, options, &success)
	return success, err
}

type SendChatActionOptions struct {
	ChatId int    `json:"chat_id"`
	Action string `json:"action"`
}

func (c *BotClient) KickChatMember(ctx context.Context, options KickChatMemberOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiKickChatMember, options, &success)
	return success, err
}

type KickChatMemberOptions struct {
	ChatId    int `json:"chat_id"`
	UserId    int `json:"user_id"`
	UntilDate int `json:"until_date,omitempty"`
}

func (c *BotClient) UnbanChatMember(ctx context.Context, options UnbanChatMemberOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiUnbanChatMember, options, &success)
	return success, err
}

type UnbanChatMemberOptions struct {
	ChatId       int  `json:"chat_id"`
	UserId       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

func (c *BotClient) RestrictChatMember(ctx context.Context, options RestrictChatMemberOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiRestrictChatMember, options, &success)
	return success, err
}

type RestrictChatMemberOptions struct {
	ChatId      int             `json:"chat_id"`
	UserId      int             `json:"user_id"`
	Permissions ChatPermissions `json:"permissions"`
	UntilDate   int             `json:"until_date,omitempty"`
}

func (c *BotClient) PromoteChatMember(ctx context.Context, options PromoteChatMemberOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiPromoteChatMember, options, &success)
	return success, err
}

type PromoteChatMemberOptions struct {
	ChatId             int  `json:"chat_id"`
	UserId             int  `json:"user_id"`
	IsAnonymous        bool `json:"is_anonymous,omitempty"`
	CanChangeInfo      bool `json:"can_change_info,omitempty"`
	CanPostMessages    bool `json:"can_post_messages,omitempty"`
	CanEditMessages    bool `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool `json:"can_delete_messages,omitempty"`
	CanInviteUsers     bool `json:"can_invite_users,omitempty"`
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	CanPinMessages     bool `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  bool `json:"can_promote_members,omitempty"`
}

func (c *BotClient) SetChatAdministratorCustomTitle(ctx context.Context, options SetChatAdministratorCustomTitleOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetChatAdministratorCustomTitle, options, &success)
	return success, err
}

type SetChatAdministratorCustomTitleOptions struct {
	ChatId      int    `json:"chat_id"`
	UserId      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func (c *BotClient) SetChatPermissions(ctx context.Context, options SetChatPermissionsOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetChatPermissions, options, &success)
	return success, err
}

type SetChatPermissionsOptions struct {
	ChatId      int             `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (c *BotClient) ExportChatInviteLink(ctx context.Context, options ChatOptions) (string, error) {
	var inviteLink string
	err := c.postJson(ctx, apiExportChatInviteLink, options, &inviteLink)
	return inviteLink, err
}

func (c *BotClient) DeleteChatPhoto(ctx context.Context, options ChatOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiDeleteChatPhoto, options, &success)
	return success, err
}

func (c *BotClient) SetChatTitle(ctx context.Context, options SetChatTitleOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetChatTitle, options, &success)
	return success, err
}

type SetChatTitleOptions struct {
	ChatId int    `json:"chat_id"`
	Title  string `json:"title"`
}

func (c *BotClient) SetChatDescription(ctx context.Context, options SetChatDescriptionOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetChatDescription, options, &success)
	return success, err
}

type SetChatDescriptionOptions struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

func (c *BotClient) PinChatMessage(ctx context.Context, options PinChatMessageOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiPinChatMessage, options, &success)
	return success, err
}

type PinChatMessageOptions struct {
	ChatId              int  `json:"chat_id"`
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

func (c *BotClient) UnpinChatMessage(ctx context.Context, options UnpinChatMessageOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiUnpinChatMessage, options, &success)
	return success, err
}

type UnpinChatMessageOptions struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id,omitempty"`
}

func (c *BotClient) UnpinAllChatMessages(ctx context.Context, options ChatOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiUnpinAllChatMessages, options, &success)
	return success, err
}

func (c *BotClient) LeaveChat(ctx context.Context, options ChatOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiLeaveChat, options, &success)
	return success, err
}

func (c *BotClient) GetChat(ctx context.Context, options ChatOptions) (*Chat, error) {
	var chat Chat
	err := c.postJson(ctx, apiGetChat, options, &chat)
	return &chat, err
}

func (c *BotClient) GetChatAdministrators(ctx context.Context, options ChatOptions) ([]ChatMember, error) {
	var chatMembers []ChatMember
	err := c.postJson(ctx, apiGetChatAdministrators, options, &chatMembers)
	return chatMembers, err
}

func (c *BotClient) GetChatMembersCount(ctx context.Context, options ChatOptions) (int, error) {
	var membersCount int
	err := c.postJson(ctx, apiGetChatMembersCount, options, &membersCount)
	return membersCount, err
}

func (c *BotClient) GetChatMember(ctx context.Context, options GetChatMemberOptions) (*ChatMember, error) {
	var member ChatMember
	err := c.postJson(ctx, apiGetChatMember, options, &member)
	return &member, err
}

type GetChatMemberOptions struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (c *BotClient) SetChatStickerSet(ctx context.Context, options SetChatStickerSetOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetChatStickerSet, options, &success)
	return success, err
}

type SetChatStickerSetOptions struct {
	ChatId         int    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func (c *BotClient) DeleteChatStickerSet(ctx context.Context, options ChatOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiDeleteChatStickerSet, options, &success)
	return success, err
}
