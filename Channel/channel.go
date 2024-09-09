package channel

/*
Channel will be the connection point for the clients. It will be responsible for managing the clients and their messages.
It will be responsible for broadcasting messages to all clients.
Keeping track of the clients in the channel.
Keeping track of the messages in the channel.
Storing the messages in a database.(Optional)
Loading the messages from the database.(Optional)

*/
import (
	client "github.com/CVilledieu/ChatServer/client"
)

type Channel struct {
	clients    []client.Client
	MessageLog MessageLog
}

type MessageLog struct {
	MostRecentMessage Message
	AmounOfMessages   uint32
}

type Message struct {
	MessageID       uint32
	Message         []byte
	PreviousMessage *Message
}

func newChannel() *Channel {
	return &Channel{
		clients: make([]client.Client, 0),
		MessageLog: MessageLog{
			AmounOfMessages: 0,
		},
	}
}

func (c *Channel) addClient(client client.Client) uint32 {
	c.clients = append(c.clients, client)
	return uint32(len(c.clients) - 1)
}

func (c *Channel) removeClient(client client.Client) {
	clientSessionId := client.getClientSessionId()
	c.clients = append(c.clients[:clientSessionId], c.clients[clientSessionId+1:]...)

}

func (c *Channel) newMessage(message []byte) {
	c.MessageLog.UpdateMessageLog(message)
}

func (m *MessageLog) UpdateMessageLog(message []byte) {
	newMessage := Message{
		MessageID:       m.MostRecentMessage.MessageID + 1,
		Message:         message,
		PreviousMessage: &m.MostRecentMessage,
	}

	m.MostRecentMessage = newMessage
	m.AmounOfMessages++
}
