package auth

import (
	"github.com/PaluMacil/walkabout/message"
	"net"
)

func DoLogin(handler message.Handler, conn net.Conn, request message.LoginRequest) {
	var session message.Session
	// for demo purposes, auth always succeeds
	if true {
		session = message.SessionFor(conn)
		session.Authenticated = true
	}
	response := message.LoginResponse{
		EntityID:  session.EntityID,
		SessionID: session.SessionID,
	}
	handler.Send(response.Envelope(conn))
}
