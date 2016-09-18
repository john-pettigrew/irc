package irc

import (
	"github.com/stretchr/testify/mock"
	"testing"
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

}
