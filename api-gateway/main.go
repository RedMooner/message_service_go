package main

import (
	"log"
	"net/http"
    "api-gateway/proxer"
)

func main() {

	messageService := proxer.MServices{EndPoints: []string{"/message"}}
	messageService.SetURL("http://message-service", 8081)
	messageService.InitializeService()

	timeService := proxer.MServices{EndPoints: []string{"/time"}}
	timeService.SetURL("http://time-service", 8082)
	timeService.InitializeService()

	log.Println("API Gateway is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
    
}

// Проксирование запросов к message-service
// messageServiceURL, _ := url.Parse("http://message-service:8081")
// messageServiceProxy := httputil.NewSingleHostReverseProxy(messageServiceURL)

// timeServiceURL, _ := url.Parse("http://time-service:8082")
// timeServiceProxy := httputil.NewSingleHostReverseProxy(timeServiceURL)

// http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request){
//     timeServiceProxy.ServeHTTP(w,r)
// })