package main

type Room struct {
	Clients    map[*User]bool
	MessageLog MessageLog
}

type MessageLog struct {
	MostRecent   MessageNode
	MessageCount uint32
}

type Message []byte

type MessageNode struct {
	Data Message
	Next *MessageNode
}

func newMessageNode(data []byte) *MessageNode {
	return &MessageNode{Data: data}
}

func (m *MessageLog) addNewNode(n *MessageNode) {
	previousHead := m.MostRecent
	n.Next = &previousHead
	m.MostRecent = *n
}

// based on distance from head.
// Zero meaning the most recent message. 1 meaning the first most recent message after the head....
func (m *MessageLog) getMessageNode(n int) *MessageNode {
	target := m.MostRecent
	for i := 0; i < n; i++ {
		target = *target.Next
	}
	return &target
}
