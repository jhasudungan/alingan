package middleware

import (
	"alingan/model"
	"errors"
	"net/http"
	"time"
)

type AuthMiddleware struct {
	SessionList map[string]*model.Session
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

	if ownerSession.Expiry.After(ownerSession.Expiry.Add(2 * time.Minute)) {
		delete(a.SessionList, sessionToken)
		return false, errors.New("authentication error - session is expired"), nil
	}

	if ownerSession.Role != "owner" {
		delete(a.SessionList, sessionToken)
		return false, errors.New("authentication error - role not permiteed"), nil
	}

	return true, nil, ownerSession
}