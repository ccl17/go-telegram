package telegram

import (
	"context"
	"net/http"
)

const (
	WEBAPIURLFORMAT = "https://api.telegram.org/bot%s/%s"
)

type BotClient struct {
	botToken   string
	httpClient *http.Client
	httpMux    *http.ServeMux
}

type Option func(*BotClient)

func OptionHttpClient(httpClient *http.Client) func(*BotClient) {
	return func(c *BotClient) {
		c.httpClient = httpClient
	}
}

func OptionHttpMux(httpMux *http.ServeMux) func(*BotClient) {
	return func(c *BotClient) {
		c.httpMux = httpMux
	}
}

func New(botToken string, options ...Option) *BotClient {
	t := &BotClient{
		botToken:   botToken,
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(t)
	}

	return t
}

func (c *BotClient) GetMe() (*Bot, error) {
	var bot Bot
	apiResp, err := doGet(context.Background(), c.httpClient, c.buildEndpoint("getMe"), &bot)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &bot, nil
}
