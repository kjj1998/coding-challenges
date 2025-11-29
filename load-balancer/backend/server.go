package main

import (
	"fmt"
	"net/http"
	"time"
)

type BackendServer struct {
	http.Handler
}

func NewBackendServer() *BackendServer {
	s := new(BackendServer)

	router := http.NewServeMux()
	router.Handle("/healthcheck", http.HandlerFunc(s.healthcheckHandler))
	router.Handle("/", http.HandlerFunc(s.helloHandler))

	s.Handler = router

	return s
}

func (*BackendServer) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received healthcheck from load balancer %v\n", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (*BackendServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v\n", r.RemoteAddr)
	fmt.Printf("%v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("User-Agent: %v\n", r.UserAgent())
	fmt.Printf("Accept: %v\n", r.Header["Accept"])
	fmt.Printf("\nReplied with a hello message\n\n")

	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Hello From Backend Server %v", r.Host)
}
