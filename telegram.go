package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

const (
	// Getting Updates
	apiGetUpdates     = "/getUpdates"
	apiSetWebhook     = "/setWebhook"
	apiDeleteWebhook  = "/deleteWebhook"
	apiGetWebhookInfo = "/getWebhookInfo"

	// Methods
	apiGetMe                           = "/getMe"
	apiLogOut                          = "/logOut"
	apiClose                           = "/close"
	apiSendMessage                     = "/sendMessage"
	apiForwardMessage                  = "/forwardMessage"
	apiCopyMessage                     = "/copyMessage"
	apiSendPhoto                       = "/sendPhoto"
	apiSendAudio                       = "/sendAudio"
	apiSendDocument                    = "/sendDocument"
	apiSendVideo                       = "/sendVideo"
	apiSendAnimation                   = "/sendAnimation"
	apiSendVoice                       = "/sendVoice"
	apiSendVideoNote                   = "/sendVideoNote"
	apiSendMediaGroup                  = "/sendMediaGroup"
	apiSendLocation                    = "/sendLocation"
	apiEditMessageLiveLocation         = "/editMessageLiveLocation"
	apiStopMessageLiveLocation         = "/stopMessageLiveLocation"
	apiSendVenue                       = "/sendVenue"
	apiSendContact                     = "/sendContact"
	apiSendPoll                        = "/sendPoll"
	apiSendDice                        = "/sendDice"
	apiSendChatAction                  = "/sendChatAction"
	apiGetUserProfilePhotos            = "/getUserProfilePhotos"
	apiGetFile                         = "/getFile"
	apiKickChatMember                  = "/kickChatMember"
	apiUnbanChatMember                 = "/unbanChatMember"
	apiRestrictChatMember              = "/restrictChatMember"
	apiPromoteChatMember               = "/promoteChatMember"
	apiSetChatAdministratorCustomTitle = "/setChatAdministratorCustomTitle"
	apiSetChatPermissions              = "/setChatPermissions"
	apiExportChatInviteLink            = "/exportChatInviteLink"
	apiSetChatPhoto                    = "/setChatPhoto"
	apiDeleteChatPhoto                 = "/deleteChatPhoto"
	apiSetChatTitle                    = "/setChatTitle"
	apiSetChatDescription              = "/setChatDescription"
	apiPinChatMessage                  = "/pinChatMessage"
	apiUnpinChatMessage                = "/unpinChatMessage"
	apiUnpinAllChatMessages            = "/unpinAllChatMessages"
	apiLeaveChat                       = "/leaveChat"
	apiGetChat                         = "/getChat"
	apiGetChatAdministrators           = "/getChatAdministrators"
	apiGetChatMembersCount             = "/getChatMembersCount"
	apiGetChatMember                   = "/getChatMember"
	apiSetChatStickerSet               = "/setChatStickerSet"
	apiDeleteChatStickerSet            = "/deleteChatStickerSet"
	apiAnswerCallbackQuery             = "/answerCallbackQuery"
	apiSetMyCommands                   = "/setMyCommands"
	apiGetMyCommands                   = "/getMyCommands"
	apiEditMessageText                 = "/editMessageText"
	apiEditMessageCaption              = "/editMessageCaption"
	apiEditMessageMedia                = "/editMessageMedia"
	apiEditMessageReplyMarkup          = "/editMessageReplyMarkup"
	apiStopPoll                        = "/stopPoll"
	apiDeleteMessage                   = "/deleteMessage"
)

type BotClient struct {
	token      string
	httpClient HttpClient
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func NewBotClient(botToken string, httpClient HttpClient) *BotClient {
	t := &BotClient{
		token:      botToken,
		httpClient: &http.Client{},
	}

	if httpClient != nil {
		t.httpClient = httpClient
	}

	return t
}

func (c *BotClient) GetMe(ctx context.Context) (*Bot, error) {
	var bot Bot
	err := c.getMethod(ctx, apiGetMe, &bot)
	return &bot, err
}

func (c *BotClient) LogOut(ctx context.Context) error {
	return c.getMethod(ctx, apiLogOut, nil)
}

func (c *BotClient) Close(ctx context.Context) error {
	return c.getMethod(ctx, apiClose, nil)
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (c *BotClient) SetMyCommands(ctx context.Context, options SetMyCommandsOptions) error {
	return c.postJson(ctx, apiSetMyCommands, options, nil)
}

type SetMyCommandsOptions struct {
	Commands []BotCommand `json:"commands"`
}

func (c *BotClient) GetMyCommands(ctx context.Context) ([]BotCommand, error) {
	var commands []BotCommand
	err := c.getMethod(ctx, apiGetMyCommands, &commands)
	return commands, err
}

func (c *BotClient) getMethod(ctx context.Context, api string, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return getJson(ctx, c.httpClient, url, out)
}

func (c *BotClient) postJson(ctx context.Context, api string, body, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return postJson(ctx, c.httpClient, url, body, out)
}

func (c *BotClient) postMultipart(ctx context.Context, api string, in, out interface{}, multipartFiles ...*multiPartFile) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return postMultipart(ctx, c.httpClient, url, in, out, multipartFiles...)
}

type multiPartFile struct {
	inputFile *InputFile
	fieldName string
}

func getJson(ctx context.Context, client HttpClient, url string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	return do(ctx, client, req, out)
}

func postJson(ctx context.Context, client HttpClient, endpoint string, body, out interface{}) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	reqBody := bytes.NewBuffer(enc)
	req, err := http.NewRequest(http.MethodPost, endpoint, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf8")
	return do(ctx, client, req, out)
}

func postMultipart(ctx context.Context, client HttpClient, endpoint string, body, out interface{}, multipartFiles ...*multiPartFile) error {
	pr, pw := io.Pipe()
	wr := multipart.NewWriter(pw)

	go func() {
		defer func() {
			wr.Close()
		}()

		for _, multipartFile := range multipartFiles {
			ioWriter, err := wr.CreateFormFile(multipartFile.fieldName, multipartFile.inputFile.Name)
			if err != nil {
				pw.CloseWithError(err)
				return
			}

			_, err = io.Copy(ioWriter, multipartFile.inputFile.Reader)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
		}

		t := reflect.TypeOf(body)
		v := reflect.ValueOf(body)
		for i := 0; i < t.NumField(); i++ {
			name, value, err := createMultipartField(t.Field(i).Name, t.Field(i).Tag.Get("json"), v.Field(i).Interface())
			if errors.Is(err, errOmitempty) || errors.Is(err, errJsonIgnore) {
				continue
			}
			err = wr.WriteField(name, value)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		pw.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, endpoint, pr)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", wr.FormDataContentType())
	return do(ctx, client, req, out)
	//_, _ = io.Copy(os.Stdout, pr)
	//return nil
}

var errJsonIgnore = errors.New("ignored when marshalling into json")

var errOmitempty = errors.New("tagged as omitempty")

func createMultipartField(name, jsonTag string, value interface{}) (string, string, error) {
	var omitempty = strings.HasSuffix(jsonTag, ",omitempty")
	fieldName := strings.TrimSuffix(jsonTag, ",omitempty")
	if fieldName == "" {
		return "", "", fmt.Errorf("json tag is missing or empty for %s", name)
	}
	if fieldName == "-" {
		return "", "", errJsonIgnore
	}

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal %v into json", value)
	}

	fieldValue := string(valueBytes)
	if omitempty && fieldValue == "null" {
		return "", "", errOmitempty
	}
	return fieldName, fieldValue, nil
}

func do(ctx context.Context, client HttpClient, req *http.Request, out interface{}) error {
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return err
	}
	defer resp.Body.Close()
	return newResponse(resp, out)
}

func newResponse(response *http.Response, out interface{}) error {
	var apiResponse ApiResponse
	err := json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return err
	}

	if !apiResponse.Ok {
		return &ApiError{
			OriginalResponse: &apiResponse,
			RequestURL:       response.Request.URL.String(),
		}
	}

	if out == nil {
		return nil
	}

	return json.Unmarshal(apiResponse.Result, out)
}

type ApiResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

type ApiError struct {
	OriginalResponse *ApiResponse
	RequestURL       string
}

func (r *ApiError) Error() string {
	return fmt.Sprintf("%s : %d %v", r.RequestURL, r.OriginalResponse.ErrorCode, r.OriginalResponse.Description)
}

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Float64(f float64) *float64 {
	return &f
}

func Bool(b bool) *bool {
	return &b
}
