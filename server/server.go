package server

/*
Server will spawn a new Channel each time a client attempts to connect to a Channel server is not already running.
Server will be responsible for listening to inital client connections, but then handing off the client to the Channel.

*/

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
