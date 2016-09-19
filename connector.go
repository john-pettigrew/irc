package irc

import (
	"bufio"
	"net"

	"github.com/john-pettigrew/irc/message"
)

func NewClient(addr string) (client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return client{}, err
	}
	return client{conn}, nil
}

type connection interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
}

type client struct {
	conn connection
}

func (c *client) SendMessage(m message.Message) error {
	toSend := []byte(message.Marshal(m))
	_, err := c.conn.Write(toSend)
	return err
}

func (c *client) SubscribeForMessages(msgCh chan message.Message) {
	reader := bufio.NewReader(c.conn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		newMsg, err := message.Unmarshal(s)
		if err != nil {
			break
		}
		msgCh <- newMsg
	}
	close(msgCh)
}
