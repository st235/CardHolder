package server

import (
	"net/http"
	"os"
)

type Server struct {
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) *Server {
	http.HandleFunc(route, handler)
	return s
}

func (s *Server) Listen(port string) *Server {
	http.ListenAndServe(port, nil)
	return s
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

