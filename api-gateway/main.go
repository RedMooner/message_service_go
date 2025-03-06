package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // Проксирование запросов к message-service
    messageServiceURL, _ := url.Parse("http://message-service:8081")
    messageServiceProxy := httputil.NewSingleHostReverseProxy(messageServiceURL)

    http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
        messageServiceProxy.ServeHTTP(w, r)
    })

    log.Println("API Gateway is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
