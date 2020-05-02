package telegram

import (
	"context"
)

func (c *BotClient) GetUpdates(ctx context.Context, options *GetUpdatesOptions) ([]Update, error) {
	var updates []Update
	_, err := doPost(ctx, c.httpClient, c.buildEndpoint("getUpdates"), options, &updates)

	return updates, err
}

type GetUpdatesOptions struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type Update struct {
	UpdateID           int                 `json:"update_id,omitempty"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	// ChannelPost        *Message            `json:"channel_post,omitempty"`
	// EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	// ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	// PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
}
