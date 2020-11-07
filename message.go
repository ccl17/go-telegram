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

func (c *BotClient) SendMessage(ctx context.Context, options SendMessageOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiSendMessage, options, &message)
	return &message, err
}

type SendMessageOptions struct {
	ChatId                   int             `json:"chat_id"`
	Text                     string          `json:"text"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ReplyToMessageId         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) ForwardMessage(ctx context.Context, options ForwardMessageOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiForwardMessage, options, &message)
	return &message, err
}

type ForwardMessageOptions struct {
	ChatId              int  `json:"chat_id"`
	FromChatId          int  `json:"from_chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	MessageId           int  `json:"message_id"`
}

func (c *BotClient) CopyMessage(ctx context.Context, options CopyMessageOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiCopyMessage, options, &message)
	return &message, err
}

type CopyMessageOptions struct {
	ChatId                   int             `json:"chat_id"`
	FromChatId               int             `json:"from_chat_id"`
	MessageId                int             `json:"message_id"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ReplyToMessageId         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageText(ctx context.Context, options EditMessageTextOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiEditMessageText, options, &message)
	return &message, err
}

type EditMessageTextOptions struct {
	ChatId                int                  `json:"chat_id,omitempty"`
	MessageId             int                  `json:"message_id,omitempty"`
	InlineMessageId       string               `json:"inline_message_id,omitempty"`
	Text                  string               `json:"text"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	Entities              []MessageEntity      `json:"entities,omitempty"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageCaption(ctx context.Context, options EditMessageCaptionOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiEditMessageCaption, options, &message)
	return &message, err
}

type EditMessageCaptionOptions struct {
	ChatId          int                  `json:"chat_id,omitempty"`
	MessageId       int                  `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	Caption         string               `json:"caption,omitempty"`
	ParseMode       string               `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageReplyMarkup(ctx context.Context, options EditMessageReplyMarkupOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiEditMessageReplyMarkup, options, &message)
	return &message, err
}

type EditMessageReplyMarkupOptions struct {
	ChatId          int                  `json:"chat_id,omitempty"`
	MessageId       int                  `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) DeleteMessage(ctx context.Context, options DeleteMessageOptions) (bool, error) {
	var success bool
	err := c.postMethod(ctx, apiDeleteMessage, options, &success)
	return success, err
}

type DeleteMessageOptions struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (c *BotClient) SendContact(ctx context.Context, options SendContactOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiSendContact, options, &message)
	return &message, err
}

type SendContactOptions struct {
	ChatId                   int      `json:"chat_id"`
	PhoneNumber              string   `json:"phone_number"`
	FirstName                string   `json:"first_name"`
	LastName                 string   `json:"last_name"`
	Vcard                    []string `json:"vcard,omitempty"`
	DisableNotification      bool     `json:"disable_notification,omitempty"`
	ReplyToMessageId         int      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool     `json:"allow_sending_without_reply,omitempty"`
}

type Poll struct {
	Id                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionId       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionsId []int  `json:"options_id"`
}

func (c *BotClient) SendPoll(ctx context.Context, options SendPollOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiSendPoll, options, &message)
	return &message, err
}

type SendPollOptions struct {
	ChatId                   int             `json:"chat_id"`
	Question                 string          `json:"question"`
	Options                  []string        `json:"options"`
	IsAnonymous              bool            `json:"is_anonymous,omitempty"`
	Type                     string          `json:"type,omitempty"`
	AllowsMultipleAnswers    bool            `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId          int             `json:"correct_option_id,omitempty"`
	Explanation              string          `json:"explanation,omitempty"`
	ExplanationParseMode     string          `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int             `json:"open_period,omitempty"`
	CloseDate                int             `json:"close_date,omitempty"`
	IsClosed                 bool            `json:"is_closed"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ReplyToMessageId         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) StopPoll(ctx context.Context, options StopPollOptions) (*Poll, error) {
	var poll Poll
	err := c.postMethod(ctx, apiStopPoll, options, &poll)
	return &poll, err
}

type StopPollOptions struct {
	ChatId      int                  `json:"chat_id,omitempty"`
	MessageId   int                  `json:"message_id,omitempty"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

func (c *BotClient) SendDice(ctx context.Context, options SendDiceOptions) (*Message, error) {
	var message Message
	err := c.postMethod(ctx, apiSendDice, options, &message)
	return &message, err
}

type SendDiceOptions struct {
	ChatId                   int         `json:"chat_id"`
	Emoji                    string      `json:"emoji,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId         int         `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}
