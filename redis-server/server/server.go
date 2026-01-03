package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/kjj1998/coding-challenges/redis-server/parser"
)

var dictionary map[string]string

func main() {
	dictionary = make(map[string]string)
	listener, err := net.Listen("tcp", ":6395")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := readRESPMessage(reader)
		if err != nil && err != io.EOF {
			log.Printf("Read error: %v", err)
			return
		}
		if err == io.EOF {
			break
		}

		response := commandSelector(message)

		_, err = conn.Write(response)
		if err != nil {
			log.Printf("Server write error: %v", err)
		}
	}
}

func readRESPMessage(reader *bufio.Reader) ([][]byte, error) {
	typeByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	// All messages from the client must be an array
	if typeByte != '*' {
		return nil, fmt.Errorf("All messages from the client must be of type array")
	}
	array, err := parser.DeserializeArray(reader)

	results := make([][]byte, len(array))
	for i, v := range array {
		byteSlice, ok := v.([]byte)
		if !ok {
			return nil, fmt.Errorf("Cannot be converted to []byte")
		}
		results[i] = byteSlice
	}

	return results, nil
}

// func deserializeCommands(bytes []byte) ([]string, error) {
// 	if len(bytes) == 0 {
// 		return []string{}, nil
// 	}

// 	commands, err := parser.Deserialize(bytes)

// 	if err != nil {
// 		return []string{}, err
// 	}

// 	return commands, nil
// }

func commandSelector(commands [][]byte) []byte {
	if len(commands) == 0 {
		return []byte("_n\r\n")
	}

	command := strings.ToLower(string(commands[0]))

	switch command {
	case "ping":
		if len(commands) == 1 {
			return parser.SerializeSimpleString("PONG")
		} else {
			commandStrings := make([]string, len(commands)-1)
			for i, v := range commands[1:] {
				commandStrings[i] = string(v)
			}

			return parser.SerializeSimpleString(strings.Join(commandStrings, " "))
			// return []byte(strings.Join(commands[1:], " "))
		}
	// case "echo":
	// 	return []byte(strings.Join(commands[1:], " "))
	// case "set":
	// 	dictionary[commands[1]] = commands[2]
	// 	return parser.SerializeSimpleString("OK")
	// case "get":
	// 	return []byte(dictionary[commands[1]])
	default:
		return []byte("_n\r\n")
	}
}
