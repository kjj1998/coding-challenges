package main

import (
	"fmt"
	"log"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v\n", r.RemoteAddr)
	fmt.Printf("%v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("User-Agent: %v\n", r.UserAgent())
	fmt.Printf("Accept: %v\n", r.Header["Accept"])
	fmt.Printf("\nReplied with a hello message\n")

	fmt.Fprint(w, "Hello From Backend Server")
}

func main() {
	handler := http.HandlerFunc(server)
	log.Fatal(http.ListenAndServe(":6050", handler))
}
