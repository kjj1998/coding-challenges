package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// Serialize returns the serialized form of the commands given.
// The serialized form is based on the Redis Serialization Protocol (RESP)
// For e.g. the string "OK" is serialized to "+OK\r\n"
// The integer 1000 is serialized to ":1000\r\n"
func Serialize(commands []string) (bytes.Buffer, error) {
	if len(commands) == 0 {
		return bytes.Buffer{}, errors.New("No command given")
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "*%d\r\n", len(commands))

	for _, command := range commands {
		fmt.Fprintf(&buf, "$%d\r\n%s\r\n", len(command), command)
	}

	return buf, nil
}

func SerializeSimpleString(command string) []byte {
	return []byte("+" + command + "\r\n")
}

func SerializeBulkString(command []byte) []byte {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "$%d\r\n", len(command))
	buf.Write(command)
	fmt.Fprint(&buf, "\r\n")

	return buf.Bytes()
}

func SerializeInteger(intVal int) []byte {
	intString := strconv.Itoa(intVal)

	return []byte(":" + intString + "\r\n")
}

func SerializeNull() []byte {
	return []byte("_\r\n")
}

func SerializeError(command string) []byte {
	return []byte("-ERR" + " " + command + "\r\n")
}
