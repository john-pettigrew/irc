package irc

import (
	"errors"
	"testing"

	"github.com/john-pettigrew/irc/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNetConn struct {
	mock.Mock
}

func (m *MockNetConn) Read(b []byte) (n int, err error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

func (m *MockNetConn) Write(b []byte) (n int, err error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

func TestSendMessage(t *testing.T) {
	mConn := new(MockNetConn)

	irc := Client{mConn}

	type testCase struct {
		input message.Message
		b     []byte
		err   error
	}
	cases := []testCase{
		testCase{
			message.Message{Command: "join", Options: []string{"#golang"}},
			[]byte("JOIN #golang\r\n"),
			nil,
		},
		testCase{
			message.Message{Command: "ping", Options: []string{"chat.freenode.net"}},
			[]byte("PING chat.freenode.net\r\n"),
			nil,
		},
		testCase{
			message.Message{Command: "user", Options: []string{"name", "name", "irc.freenode.net", "Some Cool Name"}},
			[]byte("USER name name irc.freenode.net :Some Cool Name\r\n"),
			nil,
		},
		testCase{
			message.Message{Command: "user", Options: []string{"name", "name", "irc.freenode.net", "Some Other Cool Name"}},
			[]byte("USER name name irc.freenode.net :Some Other Cool Name\r\n"),
			errors.New("Some error"),
		},
	}

	for _, c := range cases {
		mConn.On("Write", c.b).Return(len(c.b), c.err)
		err := irc.SendMessage(c.input)
		assert.Equal(t, c.err, err)
	}

}
