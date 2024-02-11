package serverbuilder

import (
	"fmt"
	"log"
	"time"
)

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{host: host, port: port}
}

func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}

func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

func (sb *serverParam) Build() *Server {
	return &Server{param: *sb}
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Printf("builder server starting on %s:%d", s.param.host, s.param.port)
	}
	fmt.Printf("builder server starting on %s:%d\n", s.param.host, s.param.port)
	return nil
}
