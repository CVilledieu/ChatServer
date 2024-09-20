package main

/*
Server will spawn a new Channel each time a client attempts to connect to a Channel server is not already running.
Server will be responsible for listening to inital client connections, but then handing off the client to the Channel.

*/

import (
	"encoding/binary"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	PORT = ":8080"
)

type MainHub struct {
	List ListOfHubs
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var openingMessage = []byte("Please enter the server Id:")

func main() {
	mainHub := createMainHub()

	http.HandleFunc("/", mainHub.connect)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func createMainHub() *MainHub {
	return &MainHub{List: nil}
}

func (mh *MainHub) connect(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer connection.Close()

	err = connection.WriteMessage(0, openingMessage)
	if err != nil {
		log.Println(err)
		return
	}
	for {

		_, data, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		checkForHub(data)

	}
}

func checkForHub(data []byte) {
	id := binary.LittleEndian.Uint32(data)

}
