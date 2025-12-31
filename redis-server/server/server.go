package main

import (
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
	listener, err := net.Listen("tcp", ":6379")
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

	message, err := io.ReadAll(conn)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	deserializedCommands, err := deserializeCommands(message)
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(deserializedCommands) == 0 {
		log.Fatal("No commands deserialized")
	}
	fmt.Println(deserializedCommands)

	response := commandSelector(deserializedCommands)

	_, err = conn.Write(response)
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}

func deserializeCommands(bytes []byte) ([]string, error) {
	commands, err := parser.Deserialize(bytes)

	if err != nil {
		return []string{}, err
	}

	return commands, nil
}

func commandSelector(commands []string) []byte {
	command := strings.ToLower(commands[0])

	switch command {
	case "ping":
		if len(commands) == 1 {
			return []byte("PONG")
		} else {
			return []byte(strings.Join(commands[1:], " "))
		}
	case "echo":
		return []byte(strings.Join(commands[1:], " "))
	case "set":
		dictionary[commands[1]] = commands[2]
		return parser.SerializeSimpleString("OK")
	case "get":
		return []byte(dictionary[commands[1]])
	default:
		return []byte(strings.Join(commands[1:], " "))
	}
}
