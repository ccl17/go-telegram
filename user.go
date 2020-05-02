package telegram

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
