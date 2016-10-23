package irc

import (
	"bufio"
	"log"
	"net"

	"github.com/john-pettigrew/irc/message"
)

// NewClient creates a new IRC connection client.
func NewClient(addr string) (Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return Client{}, err
	}
	return Client{conn}, nil
}

type connection interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
}

// Client represents a connection to an IRC server and provides helper methods.
type Client struct {
	conn connection
}

// SendMessage sends a message to the server.
func (c *Client) SendMessage(m message.Message) error {
	toSend := []byte(message.Marshal(m))
	_, err := c.conn.Write(toSend)
	return err
}

// SubscribeForMessages waits for messages from the server and send them to msgCh.
func (c *Client) SubscribeForMessages(msgCh *chan message.Message) {
	reader := bufio.NewReader(c.conn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		newMsg, err := message.Unmarshal(s)
		if err != nil {
			log.Println(err)
			break
		}

		if newMsg.Command == "PING" {
			newMsg.Command = "PONG"
			c.SendMessage(newMsg)
			continue
		}

		*msgCh <- newMsg

	}
	close(*msgCh)
}
