package auth

import (
	"github.com/PaluMacil/walkabout/message"
	"log"
	"net"
)

func DoLogin(handler message.Handler, conn net.Conn, request message.LoginRequest) {
	var session message.Session
	// for demo purposes, auth always succeeds
	if true {
		session = message.SessionFor(conn)
		session.Authenticated = true
		log.Println("Login success for", request.CharacterName, "with session ID", session.SessionID.String())
	}
	response := message.LoginResponse{
		EntityID:  session.EntityID,
		SessionID: session.SessionID,
	}
	handler.Send(response.Envelope(conn))
}
