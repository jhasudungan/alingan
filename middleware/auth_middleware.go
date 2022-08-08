package middleware

import (
	"alingan/model"
	"errors"
	"net/http"
	"time"
)

type AuthMiddleware struct {
	SessionList      map[string]*model.Session
	AgentSessionList map[string]*model.AgentSession
}

func (a *AuthMiddleware) AuthenticateAgent(r *http.Request) (bool, error, *model.AgentSession) {

	c, err := r.Cookie("alingan-session")

	if err != nil {
		return false, err, nil
	}

	sessionToken := c.Value

	agentSession, exist := a.AgentSessionList[sessionToken]

	if exist == false {
		return false, errors.New("authentication error - session is not recognized"), nil
	}

	if agentSession.Expiry == time.Now() {
		delete(a.AgentSessionList, sessionToken)
		return false, errors.New("authentication error - session is expired"), nil
	}

	if agentSession.Role != "agent" {
		delete(a.SessionList, sessionToken)
		return false, errors.New("authentication error - role not permiteed"), nil
	}

	return true, nil, agentSession

}

func (a *AuthMiddleware) AuthenticateOwner(r *http.Request) (bool, error, *model.Session) {

	c, err := r.Cookie("alingan-session")

	if err != nil {
		return false, err, nil
	}

	sessionToken := c.Value

	ownerSession, exist := a.SessionList[sessionToken]

	if exist == false {
		return false, errors.New("authentication error - session is not recognized"), nil
	}

	if ownerSession.Expiry == time.Now() {
		delete(a.SessionList, sessionToken)
		return false, errors.New("authentication error - session is expired"), nil
	}

	if ownerSession.Role != "owner" {
		delete(a.SessionList, sessionToken)
		return false, errors.New("authentication error - role not permiteed"), nil
	}

	return true, nil, ownerSession
}
