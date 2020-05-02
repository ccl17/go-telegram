package telegram

import (
	"context"
)

type Poll struct {
	Id                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionId       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionsId []int  `json:"options_id"`
}

func (c *BotClient) SendPoll(ctx context.Context, options SendPollOptions) (*Message, error) {
	var message Message
	_, err := doPost(ctx, c.httpClient, c.buildEndpoint("sendPoll"), options, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

type SendPollOptions struct {
	ChatId                int         `json:"chat_id"`
	Question              string      `json:"question"`
	Options               []string    `json:"options"`
	IsAnonymous           bool        `json:"is_anonymous,omitempty"`
	Type                  string      `json:"type,omitempty"`
	AllowsMultipleAnswers bool        `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId       int         `json:"correct_option_id,omitempty"`
	Explanation           string      `json:"explanation,omitempty"`
	ExplanationParseMode  string      `json:"explanation_parse_mode,omitempty"`
	OpenPeriod            int         `json:"open_period,omitempty"`
	CloseDate             int         `json:"close_date,omitempty"`
	IsClosed              bool        `json:"is_closed"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}
