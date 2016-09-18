package message

import (
	"errors"
	"strings"
)

type Message struct {
	Command string
	Options []string
}

var aliases = map[string]string{}

func Marshal(m Message) string {

	fullCmd := []string{}

	//get correct command
	cmd := strings.ToUpper(m.Command)
	if aliases[cmd] != "" {
		cmd = aliases[cmd]
	}

	//check for multi-word last option
	finalOption := ""
	if len(strings.Split(m.Options[len(m.Options)-1], " ")) > 1 {
		finalOption = ":" + m.Options[len(m.Options)-1]
		m.Options = m.Options[:len(m.Options)-1]
	}

	options := strings.Join(m.Options, " ")
	fullCmd = append(fullCmd, cmd)
	if options != "" {
		fullCmd = append(fullCmd, options)
	}
	if finalOption != "" {
		fullCmd = append(fullCmd, finalOption)
	}

	return strings.Join(fullCmd, " ") + "\r\n"
}

func Unmarshal(input string) (Message, error) {

	msg := Message{}

	if input == "" {
		return Message{}, errors.New("Input cannot be empty")
	}

	//remove ending characters
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)

	pieces := strings.Split(input, " ")
	if len(pieces) < 2 {
		return Message{}, errors.New("A command is required")
	}

	//Check for prefix
	if string(pieces[0][0]) == ":" {
		//remove extra data
		pieces = pieces[1:]
	}

	//get command
	msg.Command = pieces[0]
	if len(pieces) == 1 {
		return msg, nil
	}
	pieces = pieces[1:]

	//get options
	for pieceIndex, piece := range pieces {
		//Get any multi-word last argument
		if string(piece[0]) == ":" {
			pieces[pieceIndex] = string(piece[1:])
			msg.Options = append(msg.Options, strings.Join(pieces[pieceIndex:], " "))
			break
		}
		msg.Options = append(msg.Options, piece)
	}

	return msg, nil
}
