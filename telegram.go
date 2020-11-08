package telegram

import (
	"context"
	"net/http"
)

type BotClient struct {
	token      string
	httpClient httpClient
	httpMux    *http.ServeMux
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Option func(*BotClient)

func OptionHttpClient(httpClient httpClient) func(*BotClient) {
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
		token:      botToken,
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(t)
	}

	return t
}

func (c *BotClient) GetMe(ctx context.Context) (*Bot, error) {
	var bot Bot
	err := c.getMethod(ctx, apiGetMe, &bot)
	return &bot, err
}

func (c *BotClient) LogOut(ctx context.Context) (*Bot, error) {
	var bot Bot
	err := c.getMethod(ctx, apiLogOut, &bot)
	return &bot, err
}

func (c *BotClient) Close(ctx context.Context) (bool, error) {
	var success bool
	err := c.getMethod(ctx, apiClose, &success)
	return success, err
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (c *BotClient) SetMyCommands(ctx context.Context, options SetMyCommandsOptions) (bool, error) {
	var success bool
	err := c.postJson(ctx, apiSetMyCommands, options, &success)
	return success, err
}

type SetMyCommandsOptions struct {
	Commands []BotCommand `json:"commands"`
}

func (c *BotClient) GetMyCommands(ctx context.Context) ([]BotCommand, error) {
	var commands []BotCommand
	err := c.getMethod(ctx, apiGetMyCommands, &commands)
	return commands, err
}
