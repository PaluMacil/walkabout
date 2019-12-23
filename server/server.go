package main

import (
	"github.com/PaluMacil/walkabout/message"
	"github.com/PaluMacil/walkabout/message/wio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	listener       net.Listener
	messageHandler message.Handler
	queue          chan message.Envelope
}

func StartServer(port string) *Server {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	server := &Server{
		listener: listener,
		queue:    make(chan message.Envelope, 1000),
	}
	server.messageHandler = wio.TCPHandler{
		Server: server,
	}
	go server.sendMail()
	return server
}

func (s *Server) sendMail() {
	for e := range s.queue {
		e.Conn.Write(e.Raw.Bytes())
	}
}

func (s *Server) HandleSessionConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.messageHandler.HandleConnection(conn)
	}
}

func (s *Server) WaitForInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Server shutting down...")
		os.Exit(1)
	}()
}

func (s *Server) EnqueueMessage(envelope message.Envelope) {
	s.queue <- envelope
}
