package serveroptional

import (
	"fmt"
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

// Option is a functional option for configuring a server.
type Option func(*Server)

// WithXXX is a function that returns server with a specific option.
func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func New(host string, port int, options ...Option) *Server {
	server := &Server{host: host, port: port}
	for _, option := range options {
		option(server)
	}
	return server
}

func (s *Server) Start() error {
	if s.logger != nil {
		s.logger.Printf("optional server started at %s:%d\n", s.host, s.port)
	}
	fmt.Printf("optional server started at %s:%d\n", s.host, s.port)
	return nil
}
