package telegram

import (
	"context"
)

type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	Location *Location `json:"location,omitempty"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
	Location        *Location `json:"location"`
}

type InlineQueryResultArticle struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 string               `json:"url,omitempty"`
	HideUrl             bool                 `json:"hide_url,omitempty"`
	Description         string               `json:"description,omitempty"`
	ThumbUrl            string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                  `json:"thumb_width,omitempty"`
	ThumbHeight         int                  `json:"thumb_height,omitempty"`
}

type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

func (c *BotClient) InlineMode() (bool, error) {
	user, err := c.GetMe()
	if err != nil {
		return false, err
	}

	return user.IsBot && user.SupportsInlineQueries, nil
}

func (c *BotClient) AnswerInlineQuery(opt AnswerCallbackQueryOptions) (bool, error) {
	var success bool
	apiResp, err := doPost(context.Background(), c.httpClient, c.buildEndpoint("answerInlineQuery"), opt, &success)
	if err != nil {
		return false, err
	}

	if !apiResp.Ok {
		return false, newApiRespErr(apiResp)
	}

	return success, nil
}

type AnswerInlineQueryOptions struct {
	InlineQueryId     string        `json:"inline_query_id"`
	Results           []interface{} `json:"results"`
	CacheTime         int           `json:"cache_time,omitempty"`
	IsPersonal        bool          `json:"is_personal,omitempty"`
	NextOffset        string        `json:"next_offset,omitempty"`
	SwitchPmText      string        `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string        `json:"switch_pm_parameter,omitempty"`
}
