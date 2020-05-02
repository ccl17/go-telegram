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

func (c *BotClient) PinChatMessage(options PinChatMessageOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("pinChatMessage"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type PinChatMessageOptions struct {
	ChatId              int  `json:"chat_id"`
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

func (c *BotClient) UnpinChatMessage(options ChatOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("unpinChatMessage"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

func (c *BotClient) LeaveChat(options ChatOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("leaveChat"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

func (c *BotClient) GetChat(options ChatOptions) (*Chat, error) {
	var chat Chat
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("getChat"), options, &chat)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &chat, nil
}

func (c *BotClient) GetChatAdministrators(options ChatOptions) ([]*ChatMember, error) {
	var chatMembers []*ChatMember
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("getChatAdministrators"), options, &chatMembers)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return chatMembers, nil
}

func (c *BotClient) GetChatMembersCount(options ChatOptions) (int, error) {
	var membersCount int
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("getChatMembersCount"), options, &membersCount)
	if err != nil {
		return membersCount, err
	}

	if !apiResp.Ok {
		return membersCount, newApiRespErr(apiResp)
	}

	return membersCount, nil
}

func (c *BotClient) GetChatMember(options GetChatMemberOptions) (*ChatMember, error) {
	var member ChatMember
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("getChatMember"), options, &member)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &member, nil
}

type GetChatMemberOptions struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (c *BotClient) KickChatMember(options KickChatMemberOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("kickChatMember"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type KickChatMemberOptions struct {
	ChatId    int `json:"chat_id"`
	UserId    int `json:"user_id"`
	UntilDate int `json:"until_date,omitempty"`
}

func (c *BotClient) UnbanChatMember(options UnbanChatMemberOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("unbanChatMember"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type UnbanChatMemberOptions struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (c *BotClient) RestrictChatMember(options RestrictChatMemberOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("restrictChatMember"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type RestrictChatMemberOptions struct {
	ChatId      int              `json:"chat_id"`
	UserId      int              `json:"user_id"`
	Permissions *ChatPermissions `json:"permissions"`
	UntilDate   int              `json:"until_date,omitempty"`
}

func (c *BotClient) PromoteChatMember(options PromoteChatMemberOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("PromoteChatMember"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type PromoteChatMemberOptions struct {
	ChatId             int  `json:"chat_id"`
	UserId             int  `json:"user_id"`
	CanChangeInfo      bool `json:"can_change_info,omitempty"`
	CanPostMessages    bool `json:"can_post_messages,omitempty"`
	CanEditMessages    bool `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool `json:"can_delete_messages,omitempty"`
	CanInviteUsers     bool `json:"can_invite_users,omitempty"`
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	CanPinMessages     bool `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  bool `json:"can_promote_members,omitempty"`
}

func (c *BotClient) SetChatAdministratorCustomTitle(options ChatAdministratorCustomTitleOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("SetChatAdministratorCustomTitle"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type ChatAdministratorCustomTitleOptions struct {
	ChatId      int    `json:"chat_id"`
	UserId      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func (c *BotClient) SetChatPermissions(options SetChatPermissionsOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("setChatPermissions"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type SetChatPermissionsOptions struct {
	ChatId      int              `json:"chat_id"`
	Permissions *ChatPermissions `json:"permissions"`
}

// func (c *BotClient) ExportChatInviteLink(options ChatOptions) (string, error) {
// 	var inviteLink string
// 	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("exportChatInviteLink"), options, &inviteLink)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	if !apiResp.Ok {
// 		return "", newApiRespErr(apiResp)
// 	}
//
// 	return inviteLink, nil
// }

func (c *BotClient) SetChatTitle(options SetChatTitleOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("setChatTitle"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type SetChatTitleOptions struct {
	ChatId int    `json:"chat_id"`
	Title  string `json:"title"`
}

func (c *BotClient) SetChatDescription(options SetChatDescriptionOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("setChatDescription"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type SetChatDescriptionOptions struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

func (c *BotClient) SetChatStickerSet(options SetChatStickerSetOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("setChatStickerSet"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type SetChatStickerSetOptions struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (c *BotClient) DeleteChatStickerSet(options ChatOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("deleteChatStickerSet"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}
