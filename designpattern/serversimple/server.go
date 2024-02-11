package serversimple

import (
	"fmt"
)

type Server struct {
	host string
	port int
}

func New(host string, port int) *Server {
	return &Server{host: host, port: port}
}

func (s *Server) Start() error {
	fmt.Printf("simple server started at %s:%d\n", s.host, s.port)
	return nil
}
