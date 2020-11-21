package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const TEST_TOKEN = "mock-bot-token"

func setup() (botClient *BotClient, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()
	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	b := &BotClient{
		token:      TEST_TOKEN,
		httpClient: &http.Client{},
		BaseURL:    server.URL,
	}

	return b, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if r.Method != want {
		t.Errorf("Request method: %v; want %v", r.Method, want)
	}
}

func testBody(t *testing.T, r *http.Request, v interface{}) {
	t.Helper()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		t.Errorf("Request body decode error %v", err)
	}
}

func TestNewBotClient(t *testing.T) {
	b, err := NewBotClient(TEST_TOKEN, nil, "")
	if err != nil {
		t.Errorf("NewBotClient err is %v; want nil", err)
	}

	if got, want := b.BaseURL, BaseURL+b.token; got != want {
		t.Errorf("NewBotClient BaseURL is %v; want %v", got, want)
	}

	if b.httpClient == nil {
		t.Errorf("NewBotClient httpClient is nil; want default http client")
	}

	if _, err := NewBotClient("", nil, ""); err == nil {
		t.Errorf("NewBotClient err is nil; want botToken cannot be empty")
	}

}

func TestBotClient_GetMethod(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	type mockObject struct {
		FieldOne string `json:"field_one,omitempty"`
		FieldTwo int    `json:"field_two,omitempty"`
	}
	var o mockObject

	mux.HandleFunc("/mock/get", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		fmt.Fprint(w,
			`{
					"ok": true,
					"Result": {
								"field_one": "One",
								"field_two": 2
					}
				}`)
	})

	if err := b.getMethod(context.Background(), "/mock/get", &o); err != nil {
		t.Errorf("getMethod returned error %v", err)
	}

	if got := o.FieldOne; got != "One" {
		t.Errorf("getMethod FieldOne want %v; got %v", "One", got)
	}

	if got := o.FieldTwo; got != 2 {
		t.Errorf("getMethod FieldTwo want %v; got %v", 2, got)
	}
}

func TestBotClient_PostMethod(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	type mockPayload struct {
		StringField    string   `json:"string_field,omitempty"`
		IntField       int      `json:"int_field,omitempty"`
		ArrStringField []string `json:"arr_string_field,omitempty"`
	}
	p := &mockPayload{"myStringField", 555, []string{"my", "arr", "string", "field"}}

	mux.HandleFunc("/mock/post/json", func(w http.ResponseWriter, r *http.Request) {
		v := new(mockPayload)
		testMethod(t, r, http.MethodPost)
		testBody(t, r, v)

		if !reflect.DeepEqual(v, p) {
			t.Errorf("Request body = %+v, want %+v", v, p)
		}
		fmt.Fprint(w, `{"ok": true}`)
	})

	err := b.postJson(context.Background(), "/mock/post/json", p, nil)
	if err != nil {
		t.Errorf("postJson returned error %v", err)
	}
}

func TestBotClient_GetMe(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/getMe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w,
			`{
    					"ok": true,
    					"result": {
        					"id": 123456789,
        					"is_bot": true,
        					"first_name": "test_bot",
        					"username": "test_bot_username",
        					"can_join_groups": true,
        					"can_read_all_group_messages": false,
        					"supports_inline_queries": false
						}
					}`,
		)
	})

	want := &Bot{
		User: User{
			Id:        123456789,
			IsBot:     true,
			FirstName: "test_bot",
			Username:  "test_bot_username",
		},
		CanJoinGroups:           true,
		CanReadAllGroupMessages: false,
		SupportsInlineQueries:   false,
	}

	got, err := b.GetMe(context.Background())
	if err != nil {
		t.Errorf("getMe returned error %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("getMe returned %v; want %v", got, want)
	}
}

func TestBotClient_LogOut(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/logOut", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w,
			`{
    					"ok": true,
    					"result": true
					}`,
		)
	})

	if err := b.LogOut(context.Background()); err != nil {
		t.Errorf("logout returned error  %v", err)
	}
}

func TestBotClient_Close(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w,
			`{
    					"ok": true,
    					"result": true
					}`,
		)
	})

	if err := b.Close(context.Background()); err != nil {
		t.Errorf("close returned error  %v", err)
	}
}

func TestBotClient_SetMyCommands(t *testing.T) {
	b, mux, teardown := setup()
	defer teardown()

	opts := SetMyCommandsOptions{
		Commands: []BotCommand{
			{"/firstCommand", "Description of first bot command"},
			{"/secondCommand", "Description of second bot command"}},
	}

	mux.HandleFunc("/setMyCommands", func(w http.ResponseWriter, r *http.Request) {
		v := new(SetMyCommandsOptions)
		testMethod(t, r, http.MethodPost)
		testBody(t, r, v)

		if !reflect.DeepEqual(*v, opts) {
			t.Errorf("Request body = %+v, want %+v", *v, opts)
		}

		fmt.Fprintf(w,
			`{
    					"ok": true,
    					"result": true
					}`,
		)
	})

	if err := b.SetMyCommands(context.Background(), opts); err != nil {
		t.Errorf("close returned error  %v", err)
	}
}
