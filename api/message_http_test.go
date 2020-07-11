// +build message_api

package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/rinosukmandityo/message-flow/api"
	m "github.com/rinosukmandityo/message-flow/models"
	repo "github.com/rinosukmandityo/message-flow/repositories"
	rh "github.com/rinosukmandityo/message-flow/repositories/helper"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	messageRepo repo.MessageRepository
	r           *mux.Router
	ts          *httptest.Server
)

func MessageTestData() []m.Message {
	return []m.Message{
		{Message: "Message 01"},
		{Message: "Message 02"},
		{Message: "Message 03"},
	}

}

func init() {
	messageRepo = rh.ChooseRepo()
	r = RegisterHandler()
}

func TestMessageHTTP(t *testing.T) {
	ts = httptest.NewServer(r)
	defer ts.Close()

	t.Run("Insert Message", InsertMessage)
	t.Run("Get Message", GetAllMessage)
	t.Run("Websocket Message", WebsocketMessage)
}

func readMessageData(resp *http.Response) (*m.Message, error) {
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	message, _ := GetSerializer(ContentTypeJson).Decode(body)
	return message, nil
}

func PostData(t *testing.T, ts *httptest.Server, url string, _data m.Message) error {
	dataBytes, e := getBytes(_data)
	if e != nil {
		return e
	}
	resp, _, e := makeRequest(t, ts, "POST", url, bytes.NewReader(dataBytes))
	if e != nil {
		return e
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New("status should be 'Status Created' (201)")
	}

	return nil
}

func GetData(t *testing.T, ts *httptest.Server, url, expected string) error {
	resp, body, e := makeRequest(t, ts, "GET", url, nil)
	if e != nil {
		return e
	}
	if resp.StatusCode != http.StatusFound && strings.Contains(body, expected) {
		return errors.New("status should be 'Status Found' (302)")
	}

	return nil
}

func getBytes(_data m.Message) ([]byte, error) {
	dataBytes, e := json.Marshal(&_data)
	if e != nil {
		return dataBytes, e
	}
	return dataBytes, nil
}

func InsertMessage(t *testing.T) {
	t.Run("Case 1: Save data", func(t *testing.T) {
		for _, _data := range MessageTestData() {
			if e := PostData(t, ts, "/message", _data); e != nil {
				t.Errorf("[ERROR] - Failed to save data %s ", e.Error())
			}
		}
	})
}

func GetAllMessage(t *testing.T) {
	t.Run("Case 1: Get Data", func(t *testing.T) {
		_data := MessageTestData()[0]
		if e := GetData(t, ts, fmt.Sprintf("/message"), _data.Message); e != nil {
			t.Errorf("[ERROR] - Failed to get data %s", e.Error())
		}
	})
}

func WebsocketMessage(t *testing.T) {
	// Convert http://127.0.0.1 to ws://127.0.0.
	u := fmt.Sprintf("ws%s/ws", strings.TrimPrefix(ts.URL, "http"))

	// Connect to the server
	ws, _, e := websocket.DefaultDialer.Dial(u, nil)
	if e != nil {
		t.Fatalf("[ERROR] - Failed to dial websocket %s", e.Error())
	}
	defer ws.Close()

	t.Run("Case 1: Write Message", func(t *testing.T) {
		for _, data := range MessageTestData() {
			if e := ws.WriteMessage(websocket.TextMessage, []byte(data.Message)); e != nil {
				t.Fatalf("[ERROR] - Failed to write message %s", e.Error())
			}
			_, p, e := ws.ReadMessage()
			if e != nil {
				t.Errorf("[ERROR] - Failed to read message %s", e.Error())
			}
			if string(p) != data.Message {
				t.Errorf("[ERROR] - wrong message, expected (%s) got (%s)", data.Message, string(p))
			}
		}
	})
}

func makeRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string, error) {
	req, e := http.NewRequest(method, ts.URL+path, body)
	if e != nil {
		return nil, "", e
	}
	req.Header.Set("Content-Type", ContentTypeJson)

	var resp *http.Response
	switch method {
	case "GET":
		resp, e = http.DefaultTransport.RoundTrip(req)
	default:
		resp, e = http.DefaultClient.Do(req)
	}
	if e != nil {
		return nil, "", e
	}

	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, "", e
	}
	defer resp.Body.Close()

	return resp, string(respBody), nil
}
