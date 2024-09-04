package client

/*
The client package will be responsible for managing the clients.
The client will be spawned by the Channel package whenever a new client connects to the server.
The client will be responsible for sending and receiving messages.

*/
type Client struct {
	ClientID uint32
}
