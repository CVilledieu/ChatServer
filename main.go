package main

/*
Server will spawn a new Channel each time a client attempts to connect to a Channel server is not already running.
Server will be responsible for listening to inital client connections, but then handing off the client to the Channel.

*/

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	PORT = ":8080"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, r , err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		


}

