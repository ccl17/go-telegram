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
	"os"
	"path/filepath"
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

func (c *BotClient) postJson(ctx context.Context, api string, body, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return postJson(ctx, c.httpClient, url, body, out)
}

func (c *BotClient) postMultipartLocal(ctx context.Context, fullFilePath, fileFieldName, api, fieldName string, body map[string]string, out interface{}) error {
	url := "https://api.telegram.org/bot" + c.token + api
	return postMultipartLocal(ctx, c.httpClient, fullFilePath, fileFieldName, url, body, out)
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

func postMultipartLocal(ctx context.Context, client httpClient, fullFilePath, fileFieldName, endpoint string, body map[string]string, out interface{}) error {
	path, err := filepath.Abs(fullFilePath)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return postMultipart(ctx, client, endpoint, fileFieldName, filepath.Base(fullFilePath), file, body, out)
}

func postMultipart(ctx context.Context, client httpClient, endpoint, fileFieldName, fileName string, r io.Reader, body map[string]string, out interface{}) error {
	pipeReader, pipeWriter := io.Pipe()
	wr := multipart.NewWriter(pipeWriter)

	errc := make(chan error)
	go func() {
		defer pipeWriter.Close()
		ioWriter, err := wr.CreateFormFile(fileFieldName, fileName)
		if err != nil {
			errc <- err
			return
		}
		_, err = io.Copy(ioWriter, r)
		if err != nil {
			errc <- err
			return
		}

		for k, v := range body {
			err := wr.WriteField(k, v)
			if err != nil {
				errc <- err
				return
			}
		}

		if err = wr.Close(); err != nil {
			errc <- err
			return
		}
	}()

	req, err := http.NewRequest(http.MethodPost, endpoint, pipeReader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", wr.FormDataContentType())

	select {
	case err = <-errc:
		return err
	default:
		return do(ctx, client, req, out)
	}
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
