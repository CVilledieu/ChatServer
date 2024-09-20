package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type User struct {
	Name string
	Id   uint32
	Hub  *Hub
}

func startConnection(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer connection.Close()
	handleIO(connection)
}

func handleIO(connection *websocket.Conn) {
	chatLog := Log{
		MostRecent:   nil,
		FirstMessage: nil,
	}
	for {
		messageType, Data, err := connection.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		chatLog.addToLog(messageType, Data)
	}
}
