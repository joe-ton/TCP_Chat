package main

import (
    "fmt"
    "encoding/json"
    "net/http"
)

type RequestBody struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

// handle POST method
func createResource(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "POST only", http.StatusMethodNotAllowed) // 405
        return
    }

    var requestBody RequestBody
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, "Invalid JSON Payload", http.StatusBadRequest)
        return
    }

    defer r.Body.Close() // resource allocation, LIFO/stack, cleanup

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(map[string]string{
        "message": "Resource Created",
        "name": requestBody.Name,
        "email": requestBody.Email,
    })
}

func main() {
    http.HandleFunc("/resource", createResource)
    fmt.Println("Server running on:8080")
}
