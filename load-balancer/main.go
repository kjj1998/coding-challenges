package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
)

var (
	backendUrls = []string{
		"http://localhost:6050",
		"http://localhost:6060",
		"http://localhost:6070",
	}
	counter uint64
)

func server(w http.ResponseWriter, r *http.Request) {
	// Buffer all output for this requst
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Received request from %v\n", r.RemoteAddr))
	sb.WriteString(fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto))
	sb.WriteString(fmt.Sprintf("Host: %v\n", r.Host))
	sb.WriteString(fmt.Sprintf("User-Agent: %v\n", r.UserAgent()))
	sb.WriteString(fmt.Sprintf("Accept: %v\n", r.Header["Accept"]))

	index := atomic.AddUint64(&counter, 1) - 1
	backendUrl := backendUrls[index%uint64(len(backendUrls))]

	resp, err := http.Get(backendUrl)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	sb.WriteString(fmt.Sprintf("\nResponse from server: %v %v\n", resp.Proto, resp.Status))
	sb.WriteString(fmt.Sprintf("\n%s\n\n", body))

	fmt.Print(sb.String())

	fmt.Fprintf(w, "%s\n", body)
}

func main() {
	handler := http.HandlerFunc(server)
	log.Fatal(http.ListenAndServe(":6000", handler))
}
