package telegram

import (
	"context"
)

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

func (c *BotClient) AnswerCallbackQuery(ctx context.Context, options AnswerCallbackQueryOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiAnswerCallbackQuery, options, &success)
	return success, err
}

type AnswerCallbackQueryOptions struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}
