package telegram

import (
	"context"
)

type Message struct {
	MessageId            int             `json:"message_id"`
	From                 *User           `json:"from"`
	SenderChat           *Chat           `json:"sender_chat"`
	Date                 int             `json:"date"`
	Chat                 *Chat           `json:"chat"`
	ForwardFrom          *User           `json:"forward_from"`
	ForwardFromChat      *Chat           `json:"forward_from_chat"`
	ForwardFromMessageID int             `json:"forward_from_message_id"`
	ForwardSignature     string          `json:"forward_signature"`
	ForwardSenderName    string          `json:"forward_sender_name"`
	ForwardDate          int             `json:"forward_date"`
	ReplyToMessage       *Message        `json:"reply_to_message"`
	ViaBot               *Bot            `json:"via_bot"`
	EditDate             int             `json:"edit_date"`
	MediaGroupId         string          `json:"media_group_id"`
	AuthorSignature      string          `json:"author_signature"`
	Text                 string          `json:"text"`
	Entities             []MessageEntity `json:"entities"`
	Animation            *Animation      `json:"animation"`
	Audio                *Audio          `json:"audio"`
	Document             *Document       `json:"document"`
	Photo                []PhotoSize     `json:"photo"`
	//Sticker                 *Sticker                 `json:"sticker"`
	//Video                   *Video                   `json:"video"`
	VideoNote       *VideoNote      `json:"video_note"`
	Voice           *Voice          `json:"voice"`
	Caption         string          `json:"caption"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Contact         *Contact        `json:"contact"`
	Dice            *Dice           `json:"dice"`
	//Game                    *Game                    `json:"game"`
	Poll                    *Poll                        `json:"poll"`
	Venue                   *Venue                       `json:"venue"`
	Location                *Location                    `json:"location"`
	NewChatMembers          []User                       `json:"new_chat_members"`
	LeftChatMember          *User                        `json:"left_chat_member"`
	NewChatTitle            string                       `json:"new_chat_title"`
	NewChatPhoto            []PhotoSize                  `json:"new_chat_photo"`
	DeleteChatPhoto         bool                         `json:"delete_chat_photo"`
	GroupChatCreated        bool                         `json:"group_chat_created"`
	SuperGroupChatCreated   bool                         `json:"supergroup_chat_created"`
	ChannelChatCreated      bool                         `json:"channel_chat_created"`
	MigrateToChatID         int                          `json:"migrate_to_chat_id"`
	MigrateFromChatID       int                          `json:"migrate_from_chat_id"`
	PinnedMessage           *Message                     `json:"pinned_message"`
	ConnectedWebsite        string                       `json:"connected_website"`
	ProximityAlertTriggered *ProximityAlertTriggered     `json:"proximity_alert_triggered"`
	ReplyMarkup             *InlineKeyboardMarkupOptions `json:"reply_markup"`
	//Invoice               *Invoice           `json:"invoice"`
	// SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
	// PassportData          *PassportData      `json:"passport_data"`
}

type MessageId struct {
	MessageId int `json:"message_id"`
}

type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     *User  `json:"user"`
	Language string `json:"language"`
}

func (c *BotClient) SendMessage(ctx context.Context, options SendMessageOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendMessage, options, &message)
	return &message, err
}

type SendMessageOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Text                     *string         `json:"text,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    *bool           `json:"disable_web_page_preview,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) ForwardMessage(ctx context.Context, options ForwardMessageOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiForwardMessage, options, &message)
	return &message, err
}

type ForwardMessageOptions struct {
	ChatId              *int  `json:"chat_id,omitempty"`
	FromChatId          *int  `json:"from_chat_id,omitempty"`
	DisableNotification *bool `json:"disable_notification,omitempty"`
	MessageId           *int  `json:"message_id,omitempty"`
}

func (c *BotClient) CopyMessage(ctx context.Context, options CopyMessageOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiCopyMessage, options, &message)
	return &message, err
}

type CopyMessageOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	FromChatId               *int            `json:"from_chat_id,omitempty"`
	MessageId                *int            `json:"message_id,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageText(ctx context.Context, options EditMessageTextOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiEditMessageText, options, &message)
	return &message, err
}

type EditMessageTextOptions struct {
	ChatId                *int                         `json:"chat_id,omitempty"`
	MessageId             *int                         `json:"message_id,omitempty"`
	InlineMessageId       *string                      `json:"inline_message_id,omitempty"`
	Text                  *string                      `json:"text,omitempty"`
	ParseMode             *string                      `json:"parse_mode,omitempty"`
	Entities              []MessageEntity              `json:"entities,omitempty"`
	DisableWebPagePreview *bool                        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageCaption(ctx context.Context, options EditMessageCaptionOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiEditMessageCaption, options, &message)
	return &message, err
}

type EditMessageCaptionOptions struct {
	ChatId          *int                         `json:"chat_id,omitempty"`
	MessageId       *int                         `json:"message_id,omitempty"`
	InlineMessageId *string                      `json:"inline_message_id,omitempty"`
	Caption         *string                      `json:"caption,omitempty"`
	ParseMode       *string                      `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity              `json:"caption_entities,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageMedia(ctx context.Context, options EditMessageMediaOptions, media *InputFile) error {
	if media != nil {
		return c.postMultipart(ctx, apiEditMessageMedia, options, nil, &multiPartFile{media, "media"})
	}
	return c.postJson(ctx, apiEditMessageMedia, options, nil)
}

type EditMessageMediaOptions struct {
	ChatId          *int                         `json:"chat_id,omitempty"`
	MessageId       *int                         `json:"message_id,omitempty"`
	InlineMessageId *string                      `json:"inline_message_id,omitempty"`
	Media           *string                      `json:"media"`
	ReplyMarkup     *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageReplyMarkup(ctx context.Context, options EditMessageReplyMarkupOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiEditMessageReplyMarkup, options, &message)
	return &message, err
}

type EditMessageReplyMarkupOptions struct {
	ChatId          *int                         `json:"chat_id,omitempty"`
	MessageId       *int                         `json:"message_id,omitempty"`
	InlineMessageId *string                      `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) DeleteMessage(ctx context.Context, options DeleteMessageOptions) error {
	return c.postJson(ctx, apiDeleteMessage, options, nil)
}

type DeleteMessageOptions struct {
	ChatId    *int `json:"chat_id,omitempty"`
	MessageId *int `json:"message_id,omitempty"`
}

func (c *BotClient) SendContact(ctx context.Context, options SendContactOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendContact, options, &message)
	return &message, err
}

type SendContactOptions struct {
	ChatId                   *int     `json:"chat_id"`
	PhoneNumber              *string  `json:"phone_number"`
	FirstName                *string  `json:"first_name"`
	LastName                 *string  `json:"last_name,omitempty"`
	Vcard                    []string `json:"vcard,omitempty"`
	DisableNotification      *bool    `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int     `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool    `json:"allow_sending_without_reply,omitempty"`
}

type Poll struct {
	Id                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionId       int             `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`
	OpenPeriod            int             `json:"open_period"`
	CloseDate             int             `json:"close_date"`
}

type PollOption struct {
	Text       *string `json:"text,omitempty"`
	VoterCount *int    `json:"voter_count,omitempty"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionsId []int  `json:"options_id"`
}

func (c *BotClient) SendPoll(ctx context.Context, options SendPollOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendPoll, options, &message)
	return &message, err
}

type SendPollOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Question                 *string         `json:"question,omitempty"`
	Options                  []string        `json:"options,omitempty"`
	IsAnonymous              *bool           `json:"is_anonymous,omitempty"`
	Type                     *string         `json:"type,omitempty"`
	AllowsMultipleAnswers    *bool           `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId          *int            `json:"correct_option_id,omitempty"`
	Explanation              *string         `json:"explanation,omitempty"`
	ExplanationParseMode     *string         `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               *int            `json:"open_period,omitempty"`
	CloseDate                *int            `json:"close_date,omitempty"`
	IsClosed                 *bool           `json:"is_closed,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) StopPoll(ctx context.Context, options StopPollOptions) (*Poll, error) {
	var poll Poll
	err := c.postJson(ctx, apiStopPoll, options, &poll)
	return &poll, err
}

type StopPollOptions struct {
	ChatId      *int                         `json:"chat_id,omitempty"`
	MessageId   *int                         `json:"message_id,omitempty"`
	ReplyMarkup *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

func (c *BotClient) SendDice(ctx context.Context, options SendDiceOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendDice, options, &message)
	return &message, err
}

type SendDiceOptions struct {
	ChatId                   *int        `json:"chat_id,omitempty"`
	Emoji                    *string     `json:"emoji,omitempty"`
	DisableNotification      *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int        `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool       `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}
