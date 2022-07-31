package middleware

import (
	"alingan/model"
	"alingan/util"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {

	t.Run("TestAuthenticated", func(t *testing.T) {

		sessionList := make(map[string]*model.Session)

		authMiddleware := AuthMiddleware{
			SessionList: sessionList,
		}

		session := &model.Session{}
		session.Expiry = time.Now().Add(2 * time.Minute)
		session.Token = util.GenerateId("SES")
		session.Role = "owner"
		session.Id = "owner-001"

		authMiddleware.SessionList[session.Token] = session

		cookie := &http.Cookie{}
		cookie.Name = "alingan-session"
		cookie.Expires = session.Expiry
		cookie.Path = "/"
		cookie.Value = session.Token

		r, err := http.NewRequest("POST", "http://localhost:8080/test", nil)

		if err != nil {
			log.Println(err.Error())
			t.FailNow()
		}
		r.AddCookie(cookie)

		// under test
		result, err, _ := authMiddleware.AuthenticateOwner(r)

		if err != nil {
			log.Fatal(err.Error())
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestRole", func(t *testing.T) {

		sessionList := make(map[string]*model.Session)

		authMiddleware := AuthMiddleware{
			SessionList: sessionList,
		}

		session := &model.Session{}
		session.Token = util.GenerateId("SES")
		session.Expiry = time.Now().Add(2 * time.Minute)
		session.Role = "agent"
		session.Id = "agent-001"

		authMiddleware.SessionList[session.Token] = session

		cookie := &http.Cookie{}
		cookie.Name = "alingan-session"
		cookie.Expires = session.Expiry
		cookie.Path = "/"
		cookie.Value = session.Token

		r, err := http.NewRequest("POST", "http://localhost:8080/test", nil)

		if err != nil {
			log.Println(err.Error())
			t.FailNow()
		}

		r.AddCookie(cookie)

		// under test
		_, err, _ = authMiddleware.AuthenticateOwner(r)

		assert.Equal(t, "authentication error - role not permiteed", err.Error())

	})

}
