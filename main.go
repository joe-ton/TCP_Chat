
// main.go

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Define your route handlers
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/health", healthCheckHandler)

    port := ":8080"
    fmt.Printf("Server running on http://localhost%s\n", port)

    // Start the server
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    }
}

// rootHandler handles requests to "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Golang HTTP Server!")
}

// healthCheckHandler returns a simple health check status
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"status": "healthy"}`)
}

