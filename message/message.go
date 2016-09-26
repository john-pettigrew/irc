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

var optionCounts = map[string]int{
	"ADMIN":    1,
	"AWAY":     1,
	"CNOTICE":  3,
	"CPRIVMSG": 3,
	"CONNECT":  3,
	"DIE":      0,
	"ERROR":    1,
	"HELP":     0,
	"INFO":     1,
	"INVITE":   2,
	"ISON":     1,
	"JOIN":     2,
	"KICK":     3,
	"KILL":     2,
	"KNOCK":    2,
	"LINKS":    2,
	"LIST":     2,
	"LUSERS":   2,
	"MODE":     3,
	"MOTD":     1,
	"NAMES":    2,
	"NICK":     2,
	"NOTICE":   2,
	"OPER":     2,
	"PART":     2,
	"PASS":     1,
	"PING":     2,
	"PONG":     2,
	"PRIVMSG":  2,
	"QUIT":     1,
	"REHASH":   0,
	"RESTART":  0,
	"RULES":    0,
	"SERVER":   3,
	"SERVICE":  6,
	"SERVLIST": 2,
	"SQUERY":   2,
	"SQUIT":    2,
	"SETNAME":  1,
	"SILENCE":  1,
	"STATS":    2,
	"SUMMON":   3,
	"TIME":     1,
	"TOPIC":    2,
	"TRACE":    1,
	"USER":     4,
	"USERHOST": 1,
	"USERIP":   1,
	"USERS":    1,
	"VERSION":  1,
	"WALLOPS":  1,
	"WATCH":    1,
	"WHO":      2,
	"WHOIS":    2,
	"WHOWAS":   3,
}

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

func ParseCommand(input string) (Message, error) {
	msg := Message{}

	if input == "" {
		return Message{}, errors.New("Input cannot be empty")
	}

	pieces := strings.Split(input, " ")

	if pieces[0][0] == '/' {
		msg.Command = strings.ToUpper(pieces[0][1:])
		if aliases[msg.Command] != "" {
			msg.Command = aliases[msg.Command]
		}

		//remove command
		pieces = pieces[1:]
	} else {
		msg.Command = "PRIVMSG"
	}

	numOptions, cmdExists := optionCounts[msg.Command]
	if !cmdExists {
		return Message{}, errors.New("Unknown command")
	}

	for i := 0; i < numOptions; i++ {
		if i > len(pieces)-1 {
			break
		}

		if i == numOptions-1 && numOptions < len(pieces)-1 {
			msg.Options = append(msg.Options, ":"+strings.Join(pieces[i:], " "))
			break
		}

		msg.Options = append(msg.Options, pieces[i])
	}

	return msg, nil
}
