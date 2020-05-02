package telegram

import (
	"context"
)

type Message struct {
	MessageID            int              `json:"message_id,omitempty"`
	From                 *User            `json:"from,omitempty"`
	Date                 int              `json:"date,omitempty"`
	Chat                 *Chat            `json:"chat,omitempty"`
	ForwardFrom          *User            `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID int              `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string           `json:"forward_signature,omitempty"`
	ForwardSenderName    string           `json:"forward_sender_name,omitempty"`
	ForwardDate          int              `json:"forward_date,omitempty"`
	ReplyToMessage       *Message         `json:"reply_to_message,omitempty"`
	EditDate             int              `json:"edit_date,omitempty"`
	MediaGroupId         string           `json:"media_group_id,omitempty"`
	AuthorSignature      string           `json:"author_signature,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Entities             *[]MessageEntity `json:"entities,omitempty"`
	CaptionEntities      *[]MessageEntity `json:"caption_entities,omitempty"`
	// Audio                 *Audio             `json:"audio,omitempty"`
	// Document              *Document          `json:"document,omitempty"`
	// Animation             *ChatAnimation     `json:"animation,omitempty"`
	// Game                  *Game              `json:"game,omitempty"`
	// Photo                 *[]PhotoSize       `json:"photo,omitempty"`
	// Sticker               *Sticker           `json:"sticker,omitempty"`
	// Video                 *Video             `json:"video,omitempty"`
	// Voice                 *Voice             `json:"voice,omitempty"`
	// VideoNote             *VideoNote         `json:"video_note,omitempty"`
	// Caption               string             `json:"caption,omitempty"`
	// Contact               *Contact           `json:"contact,omitempty"`
	Location       *Location `json:"location,omitempty"`
	Venue          *Venue    `json:"venue,omitempty"`
	Poll           *Poll     `json:"poll,omitempty"`
	Dice           *Dice     `json:"dice,omitempty"`
	NewChatMembers *[]User   `json:"new_chat_members,omitempty"`
	LeftChatMember *User     `json:"left_chat_member,omitempty"`
	NewChatTitle   string    `json:"new_chat_title,omitempty"`
	// NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool     `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool     `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated bool     `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool     `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int      `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int      `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message `json:"pinned_message,omitempty"`
	// Invoice               *Invoice           `json:"invoice,omitempty"`
	// SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// PassportData          *PassportData      `json:"passport_data,omitempty,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   *User  `json:"user"`
}

func (c *BotClient) SendMessage(options SendMessageOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("sendMessage"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type SendMessageOptions struct {
	ChatId                int         `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

func (c *BotClient) ForwardMessage(options ForwardMessageOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("forwardMessage"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type ForwardMessageOptions struct {
	ChatId              int  `json:"chat_id"`
	FromChatId          int  `json:"from_chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	MessageId           int  `json:"message_id"`
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
	ChatId                int                   `json:"chat_id,omitempty"`
	MessageId             int                   `json:"message_id,omitempty"`
	InlineMessageId       string                `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
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
	ChatId          int                   `json:"chat_id,omitempty"`
	MessageId       int                   `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Caption         string                `json:"caption,omitempty"`
	ParseMode       string                `json:"parse_mode,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
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
	ChatId          int                   `json:"chat_id,omitempty"`
	MessageId       int                   `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) StopPoll(options StopPollOptions) (*Poll, error) {
	var poll Poll
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("stopPoll"), options, &poll)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &poll, nil
}

type StopPollOptions struct {
	ChatId      int                   `json:"chat_id,omitempty"`
	MessageId   int                   `json:"message_id,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

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

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

func (c *BotClient) SendDice(options SendDiceOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("sendDice"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type SendDiceOptions struct {
	ChatId              int         `json:"chat_id"`
	Emoji               string      `json:"emoji,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}
