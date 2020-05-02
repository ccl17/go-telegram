package telegram

import (
	"context"
)

type Message struct {
	MessageID            int              `json:"message_id, omitempty"`
	From                 *User            `json:"from, omitempty"`
	Date                 int              `json:"date, omitempty"`
	Chat                 *Chat            `json:"chat, omitempty"`
	ForwardFrom          *User            `json:"forward_from, omitempty"`
	ForwardFromChat      *Chat            `json:"forward_from_chat, omitempty"`
	ForwardFromMessageID int              `json:"forward_from_message_id, omitempty"`
	ForwardDate          int              `json:"forward_date, omitempty"`
	ReplyToMessage       *Message         `json:"reply_to_message, omitempty"`
	EditDate             int              `json:"edit_date, omitempty"`
	Text                 string           `json:"text, omitempty"`
	Entities             *[]MessageEntity `json:"entities, omitempty"`
	// CaptionEntities       *[]MessageEntity   `json:"caption_entities, omitempty"`
	// Audio                 *Audio             `json:"audio, omitempty"`
	// Document              *Document          `json:"document, omitempty"`
	// Animation             *ChatAnimation     `json:"animation, omitempty"`
	// Game                  *Game              `json:"game, omitempty"`
	// Photo                 *[]PhotoSize       `json:"photo, omitempty"`
	// Sticker               *Sticker           `json:"sticker, omitempty"`
	// Video                 *Video             `json:"video, omitempty"`
	// VideoNote             *VideoNote         `json:"video_note, omitempty"`
	// Voice                 *Voice             `json:"voice, omitempty"`
	// Caption               string             `json:"caption, omitempty"`
	// Contact               *Contact           `json:"contact, omitempty"`
	// Location              *Location          `json:"location, omitempty"`
	// Venue                 *Venue             `json:"venue, omitempty"`
	// NewChatMembers        *[]User            `json:"new_chat_members, omitempty"`
	// LeftChatMember        *User              `json:"left_chat_member, omitempty"`
	// NewChatTitle          string             `json:"new_chat_title, omitempty"`
	// NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo, omitempty"`
	// DeleteChatPhoto       bool               `json:"delete_chat_photo, omitempty"`
	// GroupChatCreated      bool               `json:"group_chat_created, omitempty"`
	// SuperGroupChatCreated bool               `json:"supergroup_chat_created, omitempty"`
	// ChannelChatCreated    bool               `json:"channel_chat_created, omitempty"`
	// MigrateToChatID       int                `json:"migrate_to_chat_id, omitempty"`
	// MigrateFromChatID     int                `json:"migrate_from_chat_id, omitempty"`
	// PinnedMessage         *Message           `json:"pinned_message, omitempty"`
	// Invoice               *Invoice           `json:"invoice, omitempty"`
	// SuccessfulPayment     *SuccessfulPayment `json:"successful_payment, omitempty"`
	// PassportData          *PassportData      `json:"passport_data,omitempty, omitempty"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   *User  `json:"user"`
}

func (c *BotClient) EditMessageText(options EditMessageTextOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("editMessageText"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type EditMessageTextOptions struct {

	// Required if inline_message_id is not specified. Unique identifier for the target
	// chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`

	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageCaption(options EditMessageCaptionOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("editMessageCaption"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type EditMessageCaptionOptions struct {

	// Required if inline_message_id is not specified. Unique identifier for the target
	// chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`

	// New text of the message, 1-4096 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// func (c *BotClient) EditMessageMedia() (*Message, error) {}

func (c *BotClient) EditMessageReplyMarkup(options EditMessageReplyMarkupOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("editMessageReplyMarkup"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type EditMessageReplyMarkupOptions struct {

	// Required if inline_message_id is not specified. Unique identifier for the target
	// chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// func (c *BotClient) StopPoll() (*Poll, error)

func (c *BotClient) DeleteMessage(options DeleteMessageOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("deleteMessage"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type DeleteMessageOptions struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}
