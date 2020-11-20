package telegram

import (
	"context"
)

type Chat struct {
	Id               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Bio              string           `json:"bio"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
	LinkedChatId     int              `json:"linked_chat_id"`
	Location         *ChatLocation    `json:"location"`
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title"`
	IsAnonymous           bool   `json:"is_anonymous"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             int    `json:"until_date"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

type ChatOptions struct {
	ChatId *int `json:"chat_id,omitempty"`
}

func (c *BotClient) SendChatAction(ctx context.Context, options SendChatActionOptions) error {
	return c.postJson(ctx, apiSendChatAction, options, nil)
}

type SendChatActionOptions struct {
	ChatId *int    `json:"chat_id,omitempty"`
	Action *string `json:"action,omitempty"`
}

func (c *BotClient) KickChatMember(ctx context.Context, options KickChatMemberOptions) error {
	return c.postJson(ctx, apiKickChatMember, options, nil)
}

type KickChatMemberOptions struct {
	ChatId    *int `json:"chat_id,omitempty"`
	UserId    *int `json:"user_id,omitempty"`
	UntilDate *int `json:"until_date,omitempty"`
}

func (c *BotClient) UnbanChatMember(ctx context.Context, options UnbanChatMemberOptions) error {
	return c.postJson(ctx, apiUnbanChatMember, options, nil)
}

type UnbanChatMemberOptions struct {
	ChatId       *int  `json:"chat_id,omitempty"`
	UserId       *int  `json:"user_id,omitempty"`
	OnlyIfBanned *bool `json:"only_if_banned,omitempty"`
}

func (c *BotClient) RestrictChatMember(ctx context.Context, options RestrictChatMemberOptions) error {
	return c.postJson(ctx, apiRestrictChatMember, options, nil)
}

type RestrictChatMemberOptions struct {
	ChatId      *int             `json:"chat_id,omitempty"`
	UserId      *int             `json:"user_id,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	UntilDate   *int             `json:"until_date,omitempty"`
}

func (c *BotClient) PromoteChatMember(ctx context.Context, options PromoteChatMemberOptions) error {
	return c.postJson(ctx, apiPromoteChatMember, options, nil)
}

type PromoteChatMemberOptions struct {
	ChatId             *int  `json:"chat_id,omitempty"`
	UserId             *int  `json:"user_id,omitempty"`
	IsAnonymous        *bool `json:"is_anonymous,omitempty"`
	CanChangeInfo      *bool `json:"can_change_info,omitempty"`
	CanPostMessages    *bool `json:"can_post_messages,omitempty"`
	CanEditMessages    *bool `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  *bool `json:"can_delete_messages,omitempty"`
	CanInviteUsers     *bool `json:"can_invite_users,omitempty"`
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`
	CanPinMessages     *bool `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  *bool `json:"can_promote_members,omitempty"`
}

func (c *BotClient) SetChatAdministratorCustomTitle(ctx context.Context, options SetChatAdministratorCustomTitleOptions) error {
	return c.postJson(ctx, apiSetChatAdministratorCustomTitle, options, nil)
}

type SetChatAdministratorCustomTitleOptions struct {
	ChatId      *int    `json:"chat_id,omitempty"`
	UserId      *int    `json:"user_id,omitempty"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

func (c *BotClient) SetChatPermissions(ctx context.Context, options SetChatPermissionsOptions) error {
	return c.postJson(ctx, apiSetChatPermissions, options, nil)
}

type SetChatPermissionsOptions struct {
	ChatId      *int             `json:"chat_id,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
}

func (c *BotClient) ExportChatInviteLink(ctx context.Context, options ChatOptions) (string, error) {
	var inviteLink string
	err := c.postJson(ctx, apiExportChatInviteLink, options, &inviteLink)
	return inviteLink, err
}

func (c *BotClient) SetChatPhoto(ctx context.Context, options SetChatPhotoOptions, photo *InputFile) error {
	if photo != nil {
		return c.postMultipart(ctx, apiSetChatPhoto, options, nil, &multiPartFile{photo, "photo"})
	}

	return c.postJson(ctx, apiSetChatPhoto, options, nil)
}

type SetChatPhotoOptions struct {
	ChatId *int    `json:"chat_id,omitempty"`
	Photo  *string `json:"photo,omitempty"`
}

func (c *BotClient) DeleteChatPhoto(ctx context.Context, options ChatOptions) error {
	return c.postJson(ctx, apiDeleteChatPhoto, options, nil)
}

func (c *BotClient) SetChatTitle(ctx context.Context, options SetChatTitleOptions) error {
	return c.postJson(ctx, apiSetChatTitle, options, nil)
}

type SetChatTitleOptions struct {
	ChatId *int    `json:"chat_id,omitempty"`
	Title  *string `json:"title,omitempty"`
}

func (c *BotClient) SetChatDescription(ctx context.Context, options SetChatDescriptionOptions) error {
	return c.postJson(ctx, apiSetChatDescription, options, nil)
}

type SetChatDescriptionOptions struct {
	ChatId      *int    `json:"chat_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (c *BotClient) PinChatMessage(ctx context.Context, options PinChatMessageOptions) error {
	return c.postJson(ctx, apiPinChatMessage, options, nil)
}

type PinChatMessageOptions struct {
	ChatId              *int  `json:"chat_id,omitempty"`
	MessageId           *int  `json:"message_id,omitempty"`
	DisableNotification *bool `json:"disable_notification,omitempty"`
}

func (c *BotClient) UnpinChatMessage(ctx context.Context, options UnpinChatMessageOptions) error {
	return c.postJson(ctx, apiUnpinChatMessage, options, nil)
}

type UnpinChatMessageOptions struct {
	ChatId    *int `json:"chat_id,omitempty"`
	MessageId *int `json:"message_id,omitempty"`
}

func (c *BotClient) UnpinAllChatMessages(ctx context.Context, options ChatOptions) error {
	return c.postJson(ctx, apiUnpinAllChatMessages, options, nil)
}

func (c *BotClient) LeaveChat(ctx context.Context, options ChatOptions) error {
	return c.postJson(ctx, apiLeaveChat, options, nil)
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
	ChatId *int `json:"chat_id,omitempty"`
	UserId *int `json:"user_id,omitempty"`
}

func (c *BotClient) SetChatStickerSet(ctx context.Context, options SetChatStickerSetOptions) error {
	return c.postJson(ctx, apiSetChatStickerSet, options, nil)
}

type SetChatStickerSetOptions struct {
	ChatId         *int    `json:"chat_id,omitempty"`
	StickerSetName *string `json:"sticker_set_name,omitempty"`
}

func (c *BotClient) DeleteChatStickerSet(ctx context.Context, options ChatOptions) error {
	return c.postJson(ctx, apiDeleteChatStickerSet, options, nil)
}
