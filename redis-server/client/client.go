package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	flag.Parse()

	var command string
	if len(flag.Args()) > 0 {
		command = strings.Join(flag.Args(), " ")
	}

	serverAddress := "localhost:6379"
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Error connecting: %v\n", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command))
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

	fmt.Print(string(response))
}
