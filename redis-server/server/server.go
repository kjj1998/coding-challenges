package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
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

	fmt.Println(string(message))

	var response string
	if string(message) == "PING" {
		response = "PONG"
	} else if strings.HasPrefix(string(message), "ECHO") {
		response = fmt.Sprintf("%s\n", strings.TrimPrefix(string(message), "ECHO "))
	} else {
		response = fmt.Sprintf("%s\n", string(message))
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
