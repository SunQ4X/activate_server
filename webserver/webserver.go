package webserver

import (
	"fmt"
	"net/http"
	"strings"
)

type WebServer struct {
	Server    *http.Server
	routerMap map[string]http.Handler
}

func NewWebServer(addr string) *WebServer {
	server := &WebServer{
		Server: &http.Server{
			Addr: addr,
		},
		routerMap: make(map[string]http.Handler),
	}

	server.Server.Handler = server

	return server
}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rcv := recover(); rcv != nil {
			fmt.Println("panic:", rcv)
		}
	}()

	handler, ok := s.routerMap[r.URL.Path]
	if ok {
		handler.ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.String(), "/static/") {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
		return
	}

}

func (s *WebServer) Handle(path string, handler http.Handler) {
	s.routerMap[path] = handler
}

func (s *WebServer) Run() {
	if err := s.Server.ListenAndServe(); err != nil {
		fmt.Println("ListenAndServe error:", err)
	}
}
