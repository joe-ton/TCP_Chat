package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("www.google.com")
    fmt.Println("resp:", resp)
    fmt.Println("err:", err)
    
}
