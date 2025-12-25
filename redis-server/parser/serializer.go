package parser

import (
	"errors"
	"fmt"
	"strconv"
)

// Serialize returns the serialized form of the commands given.
// The serialized form is based on the Redis Serialization Protocol (RESP)
// For e.g. the string "OK" is serialized to "+OK\r\n"
// The integer 1000 is serialized to ":1000\r\n"
func Serialize(commands []string) (string, error) {
	if len(commands) == 0 {
		return "", errors.New("No command given")
	}

	serializedCommand := "*" + strconv.Itoa(len(commands)) + "\r\n"

	for _, command := range commands {
		commandByteLength := len(command)

		serializedCommand = serializedCommand + fmt.Sprintf("$%d\r\n%s\r\n", commandByteLength, command)
	}

	return serializedCommand, nil
}
