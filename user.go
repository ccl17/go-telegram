package telegram

import "context"

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsBot        bool   `json:"is_bot"`
}

type Bot struct {
	User

	CanJoinGroups           bool `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool `json:"supports_inline_queries,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     []*PhotoSize `json:"photos"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

func (c *BotClient) GetUserProfilePhotos(ctx context.Context, options GetUserProfilePhotosOptions) (*UserProfilePhotos, error) {
	var userProfilePhotos UserProfilePhotos
	err := c.postJson(ctx, apiGetUserProfilePhotos, options, &userProfilePhotos)
	return &userProfilePhotos, err
}

type GetUserProfilePhotosOptions struct {
	UserId string `json:"user_id"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

func (c *BotClient) GetFile(ctx context.Context, options GetFileOptions) (*File, error) {
	var file File
	err := c.postJson(ctx, apiGetFile, options, &file)
	return &file, err
}

type GetFileOptions struct {
	FileId string `json:"file_id"`
}
