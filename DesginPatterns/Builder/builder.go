package main

import (
	"fmt"
	"time"
)

type Server struct {
	Host    string
	IP      string
	Port    int32
	TLS     bool
	TimeOut time.Duration
	MaxConn int
}

type ServerBuilder struct {
	server Server
}

func NewServerBuilder() *Server {
	return &Server{
		Host:    "server1",
		IP:      "192.168.1.1",
		Port:    80,
		TLS:     false,
		TimeOut: 2 * time.Millisecond,
		MaxConn: 10,
	}
}

func main() {

	server := NewServerBuilder()
	fmt.Println(server)

}
