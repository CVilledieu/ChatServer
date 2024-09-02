package server

func handleInput(request []byte) {
	// Do something with the request
	header, message := parseInput(request)

}

func parseInput(request []byte) (header, message []byte) {
	// Do something with the request
	return request[:64], request[64:]
}
