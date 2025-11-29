package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("port number not given")
	}

	port := flag.Args()[0]

	server := NewBackendServer()
	log.Fatal(http.ListenAndServe(":"+port, server))
}
