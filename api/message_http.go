package api

import (
	"io/ioutil"
	"net/http"

	"github.com/rinosukmandityo/message-flow/helper"
	svc "github.com/rinosukmandityo/message-flow/services"

	"github.com/pkg/errors"
)

type MessageHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	WebSocket(http.ResponseWriter, *http.Request)
}

type messageHandler struct {
	msgService svc.MessageService
}

func NewMessageHandler(msgService svc.MessageService) MessageHandler {
	return &messageHandler{msgService}
}

func (u *messageHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	data, e := u.msgService.GetAll()
	if e != nil {
		if errors.Cause(e) == helper.ErrDataNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	respBody, e := GetSerializer(contentType).Encode(data)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	SetupResponse(w, contentType, respBody, http.StatusFound)
}

func (u *messageHandler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, e := ioutil.ReadAll(r.Body)

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	msg, e := GetSerializer(contentType).Decode(requestBody)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if e = u.msgService.Store(msg); e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	SetupResponse(w, contentType, requestBody, http.StatusCreated)
}

func (u *messageHandler) WebSocket(w http.ResponseWriter, r *http.Request) {
	WSHandler(w, r)
}
