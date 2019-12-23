package wio

import (
	"bytes"
	"encoding/gob"
	"github.com/PaluMacil/walkabout/message"
	"github.com/PaluMacil/walkabout/server/auth"
	"log"
	"net"
)

type TCPHandler struct {
	Server message.Server
}

func (t TCPHandler) Send(envelope message.Envelope) {
	t.Server.EnqueueMessage(envelope)
}

func (t TCPHandler) HandleConnection(conn net.Conn) {
	session := message.SessionFor(conn)
	// receive in a new goroutine so that you can also send
	go t.receiveMessage(session)

}

func (t TCPHandler) receiveMessage(session message.Session) {
	for {
		header := message.GetHeader(session.Conn)
		messageBytes := make([]byte, header.Length)
		if _, err := session.Conn.Read(messageBytes); err != nil {
			log.Println("Failed to read from connection: ", err)
		}
		buf := bytes.NewBuffer(messageBytes)
		dec := gob.NewDecoder(buf)

		// TCPHandler message type to specific handler
		switch header.MessageType {
		case message.TypeLoginRequest:
			var m message.LoginRequest
			if err := dec.Decode(m); err != nil {
				log.Println("Could not decode LoginRequest message: ", err)
			}
			log.Println("Got LoginRequest")
			auth.DoLogin(t, session.Conn, m)
		}
	}
}
