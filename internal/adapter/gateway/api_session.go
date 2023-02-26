package gateway

import (
	"context"
	"errors"
	"sync"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
)

var _ repository.APISession = (*APISession)(nil)

type APISession struct {
	lock     sync.Mutex
	sessions map[model.SessionID]model.Session
}

func NewAPISession() *APISession {
	return &APISession{
		sessions: make(map[model.SessionID]model.Session),
	}
}

func (as *APISession) Save(ctx context.Context, session model.Session) error {
	as.lock.Lock()
	defer as.lock.Unlock()

	as.sessions[session.SessionID] = session

	return nil
}

func (as *APISession) Find(ctx context.Context, sessionID model.SessionID) (model.Session, error) {
	as.lock.Lock()
	defer as.lock.Unlock()

	if val, ok := as.sessions[sessionID]; ok {
		return val, nil
	}

	return model.Session{}, errors.New("session not found")
}

func (as *APISession) Delete(ctx context.Context, sessionID model.SessionID) error {
	as.lock.Lock()
	defer as.lock.Unlock()

	delete(as.sessions, sessionID)

	return nil
}
