package telegram

import (
	"context"
)

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareId   string    `json:"foursquare_id,omitempty"`
	FoursquareType string    `json:"foursquare_type,omitempty"`
}

func (c *BotClient) SendLocation(options SendLocationOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("sendLocation"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type SendLocationOptions struct {
	ChatId              int         `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	LivePeriod          int         `json:"live_period,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditBotMessageLiveLocation(options EditMessageLiveLocationOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("editMessageLiveLocation"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

func (c *BotClient) EditUserMessageLiveLocation(options EditMessageLiveLocationOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("editMessageLiveLocation"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type EditMessageLiveLocationOptions struct {
	ChatId          int                   `json:"chat_id,omitempty"`
	MessageId       int                   `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Latitude        float64               `json:"latitude"`
	Longitude       float64               `json:"longitude"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) StopBotMessageLiveLocation(options StopMessageLiveLocationOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("stopMessageLiveLocation"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

func (c *BotClient) StopUserMessageLiveLocation(options StopMessageLiveLocationOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("stopMessageLiveLocation"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type StopMessageLiveLocationOptions struct {
	ChatId          int                   `json:"chat_id,omitempty"`
	MessageId       int                   `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendVenue(options SendVenueOptions) (*Message, error) {
	var message Message
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("sendVenue"), options, &message)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &message, nil
}

type SendVenueOptions struct {
	ChatId              int         `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	Title               string      `json:"title"`
	Address             string      `json:"address"`
	FoursquareId        string      `json:"foursquare_id,omitempty"`
	FoursquareType      string      `json:"foursquare_type,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}
