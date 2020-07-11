package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Println(e)
	}

	reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		msgType, msg, e := conn.ReadMessage()

		if e != nil {
			log.Println(e)
			return
		}

		log.Println("New message:", string(msg))

		e = conn.WriteMessage(msgType, msg)
		if e != nil {
			log.Println(e)
			return
		}
	}
}
