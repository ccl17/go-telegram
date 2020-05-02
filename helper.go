package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *BotClient) buildEndpoint(endpoint string) string {
	return fmt.Sprintf(WEBAPIURLFORMAT, c.botToken, endpoint)
}

func doGet(ctx context.Context, client *http.Client, endpoint string, out interface{}) (*ApiResponse, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseApiResponse(resp, out)
}

func doPost(ctx context.Context, client *http.Client, endpoint string, body interface{}, out interface{}) (*ApiResponse, error) {
	req, err := requestWithPayload(endpoint, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseApiResponse(resp, out)
}

func requestWithPayload(endpoint string, payload interface{}) (*http.Request, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
}

func parseApiResponse(resp *http.Response, out interface{}) (*ApiResponse, error) {
	var apiResponse ApiResponse
	err := json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(apiResponse.Result, out); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

type ApiResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters"`
}

func newApiRespErr(apiResponse *ApiResponse) error {
	return &ApiResponseErr{apiResponse}
}

type ApiResponseErr struct {
	apiResp *ApiResponse
}

func (e *ApiResponseErr) Error() string {
	return fmt.Sprintf("Error code: %v Description: %v", e.apiResp.ErrorCode, e.apiResp.Description)
}

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}
