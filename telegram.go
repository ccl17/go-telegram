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

func (c *BotClient) GetMe(ctx context.Context) (*Bot, error) {
	var bot Bot
	apiResp, err := doGet(ctx, c.httpClient, c.buildEndpoint("getMe"), &bot)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newApiRespErr(apiResp)
	}

	return &bot, nil
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (c *BotClient) SetMyCommands(ctx context.Context, options SetMyCommandsOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(ctx, c.httpClient, c.buildEndpoint("setMyCommands"), options, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type SetMyCommandsOptions struct {
	Commands []*BotCommand `json:"commands"`
}

func (c *BotClient) GetMyCommands(ctx context.Context) ([]*BotCommand, error) {
	var commands []*BotCommand
	apiResp, err := doGet(ctx, c.httpClient, c.buildEndpoint("getMyCommands"), &commands)
	if err != nil {
		return commands, err
	}

	if !apiResp.Ok {
		return commands, newApiRespErr(apiResp)
	}

	return commands, nil
}
