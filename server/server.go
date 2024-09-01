package server

import (
	"fmt"
	"log"
	"net"
)

func Server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	_, err := conn.Read(request)
	if err != nil {
		log.Println(err)
		return
	}

	handleRequest(request)

	response := []byte("Hello, World!")
	_, err = conn.Write(response)
	if err != nil {
		log.Println(err)
		return
	}
}

func handleRequest(request []byte) {
	fmt.Println(string(request))
}
