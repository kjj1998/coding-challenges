package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/kjj1998/coding-challenges/redis-server/parser"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		serverAddress := "localhost:6379"
		conn, err := net.Dial("tcp", serverAddress)
		if err != nil {
			log.Fatalf("Error connecting: %v\n", err)
		}
		defer conn.Close()

		serializedCommands := serializeCommands(flag.Args())
		_, err = conn.Write(serializedCommands.Bytes())
		if err != nil {
			log.Fatalf("Error sending message: %v\n", err)
		}

		if tcpConn, ok := conn.(*net.TCPConn); ok {
			tcpConn.CloseWrite()
		}

		response, err := io.ReadAll(conn)
		if err != nil {
			log.Fatalf("Error reading response: %v\n", err)
		}

		deserializedResponse, err := deserialize(response)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(strings.Join(deserializedResponse, " "))
	}
}

func deserialize(bytes []byte) ([]string, error) {
	commands, err := parser.Deserialize(bytes)

	if err != nil {
		return []string{}, err
	}

	return commands, nil
}

func serializeCommands(commands []string) bytes.Buffer {
	buf, _ := parser.Serialize(commands)

	return buf
}
