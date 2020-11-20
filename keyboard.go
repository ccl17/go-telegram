package telegram

type ReplyKeyboardMarkupOptions struct {
	Keyboard        [][]KeyboardButtonOptions `json:"keyboard,omitempty"`
	ResizeKeyboard  *bool                     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool                     `json:"one_time_keyboard,omitempty"`
	Selective       *bool                     `json:"selective,omitempty"`
}

type KeyboardButtonOptions struct {
	Text            *string                        `json:"text,omitempty"`
	RequestContact  *bool                          `json:"request_contact,omitempty"`
	RequestLocation *bool                          `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollTypeOptions `json:"request_poll,omitempty"`
}

type KeyboardButtonPollTypeOptions struct {
	Type *string `json:"type,omitempty"`
}

type ReplyKeyboardRemoveOptions struct {
	RemoveKeyboard *bool `json:"remove_keyboard,omitempty"`
	Selective      *bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkupOptions struct {
	InlineKeyboard [][]InlineKeyboardButtonOptions `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButtonOptions struct {
	Text                         *string          `json:"text,omitempty"`
	Url                          *string          `json:"url,omitempty"`
	LoginUrl                     *LoginUrlOptions `json:"login_url,omitempty"`
	CallbackData                 *string          `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string          `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string          `json:"switch_inline_query_current_chat,omitempty"`
	//CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	//Pay                          bool          `json:"pay,omitempty"`
}

type LoginUrlOptions struct {
	URL                *string `json:"url,omitempty"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}
