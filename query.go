package telegram

import (
	"context"
)

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageId string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

func (c *BotClient) AnswerCallbackQuery(ctx context.Context, options AnswerCallbackQueryOptions) error {
	return c.postJson(ctx, apiAnswerCallbackQuery, options, nil)
}

type AnswerCallbackQueryOptions struct {
	CallbackQueryId *string `json:"callback_query_id,omitempty"`
	Text            *string `json:"text,omitempty"`
	ShowAlert       *bool   `json:"show_alert,omitempty"`
	Url             *string `json:"url,omitempty"`
	CacheTime       *int    `json:"cache_time,omitempty"`
}
