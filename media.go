package telegram

import (
	"context"
	"io"
)

type InputFile struct {
	io.Reader
	Name string
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
	Vcard       string `json:"vcard"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type InputMediaPhoto struct {
	Type            *string         `json:"type,omitempty"`
	Media           *string         `json:"media,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

type InputMediaVideo struct {
	Type              *string         `json:"type,omitempty"`
	Media             *string         `json:"media,omitempty"`
	Thumb             *string         `json:"thumb,omitempty"`
	Caption           *string         `json:"caption,omitempty"`
	ParseMode         *string         `json:"parse_mode,omitempty"`
	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`
	Width             *int            `json:"width,omitempty"`
	Height            *int            `json:"height,omitempty"`
	Duration          *int            `json:"duration,omitempty"`
	SupportsStreaming *bool           `json:"supports_streaming,omitempty"`
}

type InputMediaAnimation struct {
	Type            *string         `json:"type"`
	Media           *string         `json:"media"`
	Thumb           *string         `json:"thumb"`
	Caption         *string         `json:"caption"`
	ParseMode       *string         `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Width           *int            `json:"width,omitempty"`
	Height          *int            `json:"height,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
}

type InputMediaAudio struct {
	Type            *string         `json:"type,omitempty"`
	Media           *string         `json:"media,omitempty"`
	Thumb           *string         `json:"thumb,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
	Performer       *string         `json:"performer,omitempty"`
	Title           *string         `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        *string         `json:"type,omitempty"`
	Media                       *string         `json:"media,omitempty"`
	Thumb                       *string         `json:"thumb,omitempty"`
	Caption                     *string         `json:"caption,omitempty"`
	ParseMode                   *string         `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

func (c *BotClient) GetFile(ctx context.Context, options GetFileOptions) (*File, error) {
	var file File
	err := c.postJson(ctx, apiGetFile, options, &file)
	return &file, err
}

type GetFileOptions struct {
	FileId *string `json:"file_id"`
}

func (c *BotClient) SendPhoto(ctx context.Context, options SendPhotoOptions, photo *InputFile) (*Message, error) {
	var message Message

	if photo != nil {
		err := c.postMultipart(ctx, apiSendPhoto, options, &message, &multiPartFile{photo, "photo"})
		return &message, err
	}

	err := c.postJson(ctx, apiSendPhoto, options, &message)
	return &message, err
}

type SendPhotoOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Photo                    *string         `json:"photo,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendAudio(ctx context.Context, options SendAudioOptions, audio, thumb *InputFile) (*Message, error) {
	var message Message

	if audio != nil {
		multipartFiles := []*multiPartFile{{audio, "audio"}}
		if thumb != nil {
			multipartFiles = append(multipartFiles, &multiPartFile{thumb, thumb.Name})
		}
		err := c.postMultipart(ctx, apiSendAudio, options, &message, multipartFiles...)
		return &message, err
	}

	if thumb != nil {
		err := c.postMultipart(ctx, apiSendAudio, options, &message, &multiPartFile{thumb, thumb.Name})
		return &message, err
	}

	err := c.postJson(ctx, apiSendAudio, options, &message)
	return &message, err
}

type SendAudioOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Audio                    *string         `json:"audio,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	Duration                 *int            `json:"duration,omitempty"`
	Performer                *string         `json:"performer,omitempty"`
	Title                    *string         `json:"title,omitempty"`
	Thumb                    *string         `json:"thumb,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendDocument(ctx context.Context, options SendDocumentOptions, document, thumb *InputFile) (*Message, error) {
	var message Message

	if document != nil {
		multipartFiles := []*multiPartFile{{document, "document"}}
		if thumb != nil {
			multipartFiles = append(multipartFiles, &multiPartFile{thumb, thumb.Name})
		}
		err := c.postMultipart(ctx, apiSendDocument, options, &message, multipartFiles...)
		return &message, err
	}

	if thumb != nil {
		err := c.postMultipart(ctx, apiSendDocument, options, &message, &multiPartFile{thumb, thumb.Name})
		return &message, err
	}

	err := c.postJson(ctx, apiSendDocument, options, &message)
	return &message, err
}

type SendDocumentOptions struct {
	ChatId                      *int            `json:"chat_id,omitempty"`
	Document                    *string         `json:"document,omitempty"`
	Thumb                       *string         `json:"thumb,omitempty"`
	Caption                     *string         `json:"caption,omitempty"`
	ParseMode                   *string         `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool           `json:"disable_content_type_detection,omitempty"`
	DisableNotification         *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId            *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendVideo(ctx context.Context, options SendVideoOptions, video, thumb *InputFile) (*Message, error) {
	var message Message

	if video != nil {
		multipartFiles := []*multiPartFile{{video, "video"}}
		if thumb != nil {
			multipartFiles = append(multipartFiles, &multiPartFile{thumb, thumb.Name})
		}
		err := c.postMultipart(ctx, apiSendVideo, options, &message, multipartFiles...)
		return &message, err
	}

	if thumb != nil {
		err := c.postMultipart(ctx, apiSendVideo, options, &message, &multiPartFile{thumb, thumb.Name})
		return &message, err
	}

	err := c.postJson(ctx, apiSendVideo, options, &message)
	return &message, err
}

type SendVideoOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Video                    *string         `json:"video,omitempty"`
	Duration                 *int            `json:"duration,omitempty"`
	Width                    *int            `json:"width,omitempty"`
	Height                   *int            `json:"height,omitempty"`
	Thumb                    *string         `json:"thumb,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	SupportsStreaming        *bool           `json:"supports_streaming,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendAnimation(ctx context.Context, options SendAnimationOptions, animation, thumb *InputFile) (*Message, error) {
	var message Message

	if animation != nil {
		multipartFiles := []*multiPartFile{{animation, "animation"}}
		if thumb != nil {
			multipartFiles = append(multipartFiles, &multiPartFile{thumb, thumb.Name})
		}
		err := c.postMultipart(ctx, apiSendAnimation, options, &message, multipartFiles...)
		return &message, err
	}

	if thumb != nil {
		err := c.postMultipart(ctx, apiSendAnimation, options, &message, &multiPartFile{thumb, thumb.Name})
		return &message, err
	}

	err := c.postJson(ctx, apiSendAnimation, options, &message)
	return &message, err
}

type SendAnimationOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Video                    *string         `json:"video,omitempty"`
	Duration                 *int            `json:"duration,omitempty"`
	Width                    *int            `json:"width,omitempty"`
	Height                   *int            `json:"height,omitempty"`
	Thumb                    *string         `json:"thumb,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	SupportsStreaming        *bool           `json:"supports_streaming,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) sendVoice(ctx context.Context, options SendVoiceOptions, voice *InputFile) (*Message, error) {
	var message Message

	if voice != nil {
		err := c.postMultipart(ctx, apiSendVoice, options, &message, &multiPartFile{voice, "voice"})
		return &message, err
	}

	err := c.postJson(ctx, apiSendVoice, options, &message)
	return &message, err
}

type SendVoiceOptions struct {
	ChatId                   *int            `json:"chat_id,omitempty"`
	Voice                    *string         `json:"voice,omitempty"`
	Caption                  *string         `json:"caption,omitempty"`
	ParseMode                *string         `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	Duration                 *int            `json:"duration,omitempty"`
	DisableNotification      *bool           `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

func (c *BotClient) sendVideoNote(ctx context.Context, options SendVideoNoteOptions, videoNote, thumb *InputFile) (*Message, error) {
	var message Message

	if videoNote != nil {
		multipartFiles := []*multiPartFile{{videoNote, "videoNote"}}
		if thumb != nil {
			multipartFiles = append(multipartFiles, &multiPartFile{thumb, thumb.Name})
		}
		err := c.postMultipart(ctx, apiSendVideoNote, options, &message, multipartFiles...)
		return &message, err
	}

	if thumb != nil {
		err := c.postMultipart(ctx, apiSendVideoNote, options, &message, &multiPartFile{thumb, thumb.Name})
		return &message, err
	}

	err := c.postJson(ctx, apiSendVideoNote, options, &message)
	return &message, err
}

type SendVideoNoteOptions struct {
	ChatId                   *int        `json:"chat_id,omitempty"`
	VideoNote                *string     `json:"voice,omitempty"`
	Duration                 *int        `json:"duration,omitempty"`
	Length                   *int        `json:"length,omitempty"`
	Thumb                    *string     `json:"thumb,omitempty"`
	DisableNotification      *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int        `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool       `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendMediaGroup(ctx context.Context, options SendMediaGroupOptions, inputs []*InputFile) ([]Message, error) {
	var messages []Message

	if inputs == nil || len(inputs) < 1 {
		err := c.postJson(ctx, apiSendMediaGroup, options, &messages)
		return messages, err
	}

	files := make([]*multiPartFile, len(inputs))
	for index, input := range inputs {
		files[index] = &multiPartFile{input, input.Name}
	}

	err := c.postMultipart(ctx, apiSendMediaGroup, options, &messages, files...)
	return messages, err
}

type SendMediaGroupOptions struct {
	ChatId                   *int        `json:"chat_id,omitempty"`
	Media                    interface{} `json:"media,omitempty"`
	DisableNotification      *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int        `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool       `json:"allow_sending_without_reply,omitempty"`
}
