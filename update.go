package telegram

import (
	"context"
)

type Update struct {
	UpdateId           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	Poll               *Poll               `json:"poll"`
	PollAnswer         *PollAnswer         `json:"poll_answer"`
	//ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	// PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

type WebhookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	IpAddress            string   `json:"ip_address"`
	LastErrorDate        string   `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

func (c *BotClient) GetUpdates(ctx context.Context, options GetUpdatesOptions) ([]Update, error) {
	var updates []Update
	err := c.postJson(ctx, apiGetUpdates, options, &updates)
	return updates, err
}

type GetUpdatesOptions struct {
	Offset         *int     `json:"offset,omitempty"`
	Limit          *int     `json:"limit,omitempty"`
	Timeout        *int     `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (c *BotClient) SetWebhook(ctx context.Context, options SetWebhookOptions, certificate *InputFile) error {
	if certificate != nil {
		return c.postMultipart(ctx, apiSetWebhook, options, nil, &multiPartFile{certificate, "certificate"})
	}
	return c.postJson(ctx, apiSetWebhook, options, nil)
}

type SetWebhookOptions struct {
	Url                string   `json:"url,omitempty"`
	IpAddress          string   `json:"ip_address,omitempty"`
	MaxConnections     int      `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool     `json:"drop_pending_updates,omitempty"`
}

func (c *BotClient) DeleteWebhook(ctx context.Context, options DeleteWebhookOptions) error {
	return c.postJson(ctx, apiDeleteWebhook, options, nil)
}

type DeleteWebhookOptions struct {
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
}

func (c *BotClient) GetWebhookInfo(ctx context.Context) (*WebhookInfo, error) {
	var webhookInfo WebhookInfo
	err := c.getMethod(ctx, apiGetWebhookInfo, &webhookInfo)
	return &webhookInfo, err
}
