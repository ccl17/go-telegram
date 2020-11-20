package telegram

import "context"

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Bot struct {
	User
	CanJoinGroups           bool `json:"can_join_groups"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool `json:"supports_inline_queries"`
}

type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

func (c *BotClient) GetUserProfilePhotos(ctx context.Context, options GetUserProfilePhotosOptions) (*UserProfilePhotos, error) {
	var userProfilePhotos UserProfilePhotos
	err := c.postJson(ctx, apiGetUserProfilePhotos, options, &userProfilePhotos)
	return &userProfilePhotos, err
}

type GetUserProfilePhotosOptions struct {
	UserId *string `json:"user_id,omitempty"`
	Offset *int    `json:"offset,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}
