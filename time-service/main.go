package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
)

func main(){
    fmt.Println("test service")
    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet{
            now := time.Now()
            fmt.Fprintf(w, "time:", now)
        } else {
            http.Error(w,"Method not allowrd", http.StatusMethodNotAllowed)
        }

        log.Println("time service is running on port 8082")
    })
    log.Fatal(http.ListenAndServe(":8082", nil))
}