package main

import (
    "fmt"
    "errors"
    "encoding/json"
    "net/http"
)

type RequestPayload struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

// handle POST method
func createResource(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "POST only", http.StatusMethodNotAllowed)
        return
    }

    var requestPayload RequestPayload
    err := json.NewDecoder(r.Body).Decode(&requestPayload)
    if err != nil {
        http.Error(w, "Invalid Payload", http.StatusBadRequest)
    }

    defer r.Body.Close()

    w.Header().Set("Content-Type", "applciation/json")
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder
}
