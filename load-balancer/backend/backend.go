package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v\n", r.RemoteAddr)
	fmt.Printf("%v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("User-Agent: %v\n", r.UserAgent())
	fmt.Printf("Accept: %v\n", r.Header["Accept"])
	fmt.Printf("\nReplied with a hello message\n\n")

	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Hello From Backend Server %v", r.Host)
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("port number not given")
	}

	port := flag.Args()[0]

	handler := http.HandlerFunc(server)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
