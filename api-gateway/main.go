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

    timeServiceURL, _ := url.Parse("http://time-service:8082")
    timeServiceProxy := httputil.NewSingleHostReverseProxy(timeServiceURL)

    http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
        messageServiceProxy.ServeHTTP(w, r)
    })
    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request){
        timeServiceProxy.ServeHTTP(w,r)
    })

    log.Println("API Gateway is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
