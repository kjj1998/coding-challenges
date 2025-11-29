package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type backend struct {
	Address string
	Healthy bool
}

var (
	backends = []*backend{
		newBackend("6050"),
		newBackend("6060"),
		newBackend("6070"),
	}
	counter uint64
	mu      sync.RWMutex
)

// Health check goroutine
func healthCheck(interval time.Duration, healthcheck string) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for _, backend := range backends {
			go checkHealth(backend, healthcheck)
		}
	}
}

func checkHealth(backend *backend, healthcheck string) {
	// Use a short timeout for health checks
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(backend.Address + healthcheck)

	mu.Lock()
	defer mu.Unlock()

	if err != nil || resp.StatusCode != http.StatusOK {
		if backend.Healthy {
			fmt.Printf("Backend %s is DOWN\n", backend.Address)
		}
		backend.Healthy = false
	} else {
		if !backend.Healthy {
			fmt.Printf("Backend %s is UP\n", backend.Address)
		}
		backend.Healthy = true
		resp.Body.Close()
	}
}

// Get next healthy backend using round-robin
func getNextBackend() *backend {
	mu.RLock()
	defer mu.RUnlock()

	// Try each backend starting from current index
	for range backends {
		index := atomic.AddUint64(&counter, 1) - 1
		backend := backends[index%uint64(len(backends))]

		if backend.Healthy {
			return backend
		}
	}

	return nil
}

func newBackend(port string) *backend {
	s := backend{Address: "http://localhost:" + port}
	s.Healthy = true

	return &s
}

func server(w http.ResponseWriter, r *http.Request) {
	// Get the next available server
	backend := getNextBackend()
	if backend == nil {
		http.Error(w, "No healthy backends", http.StatusServiceUnavailable)
		return
	}

	// Buffer all output for this request
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Received request from %v\n", r.RemoteAddr))
	sb.WriteString(fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto))
	sb.WriteString(fmt.Sprintf("Host: %v\n", r.Host))
	sb.WriteString(fmt.Sprintf("User-Agent: %v\n", r.UserAgent()))
	sb.WriteString(fmt.Sprintf("Accept: %v\n", r.Header["Accept"]))

	resp, err := http.Get(backend.Address)
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
	healthcheck := flag.String("h", "/healthcheck", "The backend health check endpoint, defaults to /healthcheck")
	interval := flag.Int("i", 10, "The time interval (in seconds) used to check backend health, defaults to 10")

	flag.Parse()

	// start health checks in the background (check every 10 seconds)
	go healthCheck(time.Duration(*interval)*time.Second, *healthcheck)

	handler := http.HandlerFunc(server)
	fmt.Println("Load balancer running on :6000")
	log.Fatal(http.ListenAndServe(":6000", handler))
}
