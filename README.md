# IRC
An IRC client library written in Golang.

## Installing
You can install this library by running:
```
go get github.com/john-pettigrew/irc
```
Add the following to the top of your file:
```
import "github.com/john-pettigrew/irc"
```
Or if you need the message library:
```
import "github.com/john-pettigrew/irc/message"
```

## Usage
### Sample Usage
```
// create our channel and connect to the server.
msgCh := make(chan message.Message)
ircConn, err := irc.NewClient("Someserver:6667")
if err != nil {
  log.Fatal("Error connecting")
}

// Now, we can subscribe for any messages.
go ircConn.SubscribeForMessages(&msgCh)

// And, send a message.
m := message.Message{Command: "PRIVMSG", Options: []string{"Hello, server!"}}
err = ircConn.SendMessage(m)
if err != nil {
  log.Fatal("Error sending message")
}
```

### Docs
Documentation can be found [here](https://godoc.org/github.com/john-pettigrew/irc).

## Running Tests
```
go test ./...
```
