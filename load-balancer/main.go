package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v\n", r.RemoteAddr)
	fmt.Printf("%v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("User-Agent: %v\n", r.UserAgent())
	fmt.Printf("Accept: %v\n", r.Header["Accept"])

	backendUrl := "http://localhost:6050"

	resp, err := http.Get(backendUrl)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("\nResponse from server: %v %v\n", resp.Proto, resp.Status)
	fmt.Printf("\n%s\n", body)

	fmt.Fprintf(w, "%s\n", body)
}

func main() {
	handler := http.HandlerFunc(server)
	log.Fatal(http.ListenAndServe(":6000", handler))
}
