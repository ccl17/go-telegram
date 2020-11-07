package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	apiGetMe                           = "/getMe"
	apiLogOut                          = "/logOut"
	apiClose                           = "/close"
	apiSendMessage                     = "/sendMessage"
	apiForwardMessage                  = "/forwardMessage"
	apiCopyMessage                     = "/copyMessage"
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
	apiEditMessageReplyMarkup          = "/editMessageReplyMarkup"
	apiStopPoll                        = "/stopPoll"
	apiDeleteMessage                   = "/deleteMessage"
)

func (c *BotClient) getMethod(ctx context.Context, api string, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return getJson(ctx, c.httpClient, url, out)
}

func (c *BotClient) postMethod(ctx context.Context, api string, body, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return postJson(ctx, c.httpClient, url, body, out)
}

func getJson(ctx context.Context, client httpClient, url string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	return do(ctx, client, req, out)
}

func postJson(ctx context.Context, client httpClient, endpoint string, body, out interface{}) error {
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

func do(ctx context.Context, client httpClient, req *http.Request, out interface{}) error {
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return err
	}

	if !apiResponse.Ok {
		return errors.New(fmt.Sprintf("Api Error : %v", apiResponse))
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

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}
