package api

import (
	"log"
	"net"
)

type Port int16

type server struct {
	l *log.Logger
	r map[Path]handler
}

func NewServer(l *log.Logger) *server {
	return &server{l: l}
}

func (s *server) AddHandler(path Path, handler handler) {
	s.r[path] = handler
}

func (s *server) Run(port Port) {
	listen, err := net.Listen("tcp", string(port))

	if err != nil {
		s.l.Fatalf("couldn't start server: %w\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			s.l.Printf("error appeared while listening: %w\n", err)
		}

		go s.handleConnection(conn)
	}
}

func (s *server) handleConnection(conn net.Conn) {
	defer conn.Close()

	s.serveConnection(path, conn)
}

func (s *server) serveConnection(path Path, conn net.Conn) {
	h, ok := s.r[path]

	if !ok {
		s.l.Printf("no handler for this path\n")
	}

	h(conn)
}
