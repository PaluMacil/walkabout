package message

import (
	"github.com/google/uuid"
	"net"
	"sync"
	"time"
)

type sessionList struct {
	m    sync.RWMutex
	list []Session
}

func (s *sessionList) Get(sessionID uuid.UUID) (Session, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	for _, session := range s.list {
		if session.SessionID == sessionID {
			return session, nil
		}
	}
	return Session{}, SessionNotFoundError{sessionID}
}

func (s *sessionList) Add(newSession Session) {
	s.m.Lock()
	defer s.m.Unlock()
	s.list = append(s.list, newSession)
}

var Sessions = &sessionList{}

type Session struct {
	Conn          net.Conn
	Authenticated bool
	SessionID     uuid.UUID
	EntityID      uuid.UUID
	Created       time.Time
}

func SessionFor(conn net.Conn) Session {
	session := Session{
		Conn:      conn,
		SessionID: uuid.New(),
		EntityID:  uuid.New(),
		Created:   time.Now(),
	}

	Sessions.Add(session)
	return session
}
