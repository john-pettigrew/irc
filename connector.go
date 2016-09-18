package irc

import (
	"./message"
)

func NewClient(addr string) (client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return client{}, err
	}
	return client{conn}
}

var connection interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
}

type client struct {
	conn connection
}

func (i *client) SendMessage(m Message) error {
	return nil
}

func (i *client) SubscribeForMessages() {

}
