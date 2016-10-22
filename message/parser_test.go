package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	type testCase struct {
		m      Message
		result string
	}
	cases := []testCase{
		testCase{
			Message{Command: "join", Options: []string{"#golang"}},
			"JOIN #golang\r\n",
		},
		testCase{
			Message{Command: "ping", Options: []string{"chat.freenode.net"}},
			"PING chat.freenode.net\r\n",
		},
		testCase{
			Message{Command: "user", Options: []string{"name", "name", "irc.freenode.net", "Some Cool Name"}},
			"USER name name irc.freenode.net :Some Cool Name\r\n",
		},
	}

	for _, c := range cases {
		result := Marshal(c.m)
		assert.Equal(t, c.result, result)
	}
}

func TestUnmarshal(t *testing.T) {
	type testCase struct {
		input string
		m     Message
		err   error
	}

	cases := []testCase{
		testCase{
			"JOIN #golang\r\n",
			Message{Command: "JOIN", Options: []string{"#golang"}},
			nil,
		},
		testCase{
			"PING chat.freenode.net\r\n",
			Message{Command: "PING", Options: []string{"chat.freenode.net"}},
			nil,
		},
		testCase{
			"USER name name irc.freenode.net :Some Cool Name\r\n",
			Message{Command: "USER", Options: []string{"name", "name", "irc.freenode.net", "Some Cool Name"}},
			nil,
		},
		testCase{
			":aaa!bbb@ccc/ddd PRIVMSG #golang :Some cool message!",
			Message{Prefix: ":aaa!bbb@ccc/ddd", Command: "PRIVMSG", Options: []string{"#golang", "Some cool message!"}},
			nil,
		},
	}

	for _, c := range cases {
		m, err := Unmarshal(c.input)
		assert.Nil(t, c.err, err)
		assert.Equal(t, c.m, m)
	}
}
