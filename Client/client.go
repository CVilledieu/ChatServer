package client

import (
	"github.com/CVilledieu/ChatServer/channel"
	"github.com/gorilla/websocket"
)

/*
The client package will be responsible for managing the clients.
The client will be spawned by the Channel package whenever a new client connects to the server.
The client will be responsible for sending and receiving messages.
*/
type Client struct {
	ClientID        uint32
	ClientSessionID uint32
	Channel         *channel.Channel
	// The websocket connection.
	Conn *websocket.Conn
}

func (c *Client) getClientSessionId() uint32 {
	return c.ClientID

}

func (c *Client) setClientSessionId(id uint32) {
	c.ClientSessionID = id
}

func (c *Client) ConnectToChannel(channel *Channel) {
	c.Channel = channel
	c.Channel.addClient(*c)
}
