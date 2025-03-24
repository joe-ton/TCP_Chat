package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createResource(w http.ResponseWriter, r *http.Request) {
    // Guard Clauses
    // POST check
    // payload struct check
}
