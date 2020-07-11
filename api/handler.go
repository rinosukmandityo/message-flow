package api

import (
	rh "github.com/rinosukmandityo/message-flow/repositories/helper"
	"github.com/rinosukmandityo/message-flow/services/logic"

	"github.com/gorilla/mux"
)

func RegisterHandler() *mux.Router {
	r := mux.NewRouter()
	registerMessageHandler(r, NewMessageHandler(logic.NewMessageService(rh.ChooseRepo())))

	return r
}

func registerMessageHandler(r *mux.Router, handler MessageHandler) {
	r.HandleFunc("/message", handler.Post).Methods("POST")  // POST
	r.HandleFunc("/message", handler.GetAll).Methods("GET") // GET

	r.HandleFunc("/ws", handler.WebSocket)
}
