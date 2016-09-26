package irc

import (
	"bufio"
	"log"
	"net"

	"github.com/john-pettigrew/irc/message"
)

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

type Client struct {
	conn connection
}

func (c *Client) SendMessage(m message.Message) error {
	toSend := []byte(message.Marshal(m))
	_, err := c.conn.Write(toSend)
	return err
}

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

		*msgCh <- newMsg
	}
	close(*msgCh)
}
