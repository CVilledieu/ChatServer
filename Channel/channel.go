package channel

/*
Channel will be the connection point for the clients. It will be responsible for managing the clients and their messages.
It will be responsible for broadcasting messages to all clients.
Keeping track of the clients in the channel.
Keeping track of the messages in the channel.
Storing the messages in a database.(Optional)
Loading the messages from the database.(Optional)

*/
import client "github.com/CVilledieu/ChatServer/Client"

type channel struct {
	clients []client.Client
}
