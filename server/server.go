package server

import (
	"net/http"
)

func StartServer() {
	server := &http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
