package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            fmt.Fprintf(w, "Hello from Message Service1!")
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Message Service is running on port 8081...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
