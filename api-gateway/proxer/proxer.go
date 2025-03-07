package proxer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type MServices struct {
	serviceURL   *url.URL
	serviceProxy *httputil.ReverseProxy
	EndPoints    []string
}

func (s MServices) URL() string {
	return s.serviceURL.Host
}
func (s *MServices) SetURL(_url string, _port int) {
	s.serviceURL, _ = url.Parse(fmt.Sprintf("%s:%d", _url, _port))
}

func (s MServices) ServiceProxy() *httputil.ReverseProxy {
	return s.serviceProxy
}

func (s MServices) InitializeService() {
	fmt.Println("Service is starting now...")

	s.serviceProxy = httputil.NewSingleHostReverseProxy(s.serviceURL)

	for _, value := range s.EndPoints {
		fmt.Println("Handle ", value)
		http.HandleFunc(value, func(w http.ResponseWriter, r *http.Request) {
			s.serviceProxy.ServeHTTP(w, r)
		})
	}
}
