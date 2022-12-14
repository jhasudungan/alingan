package service

import (
	"alingan/model"
	"alingan/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	t.Run("TestOwnerRegistration", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.OwnerRegistrationRequest{}
		request.OwnerEmail = "admin@jfc.com"
		request.OwnerName = "Jeremi Fried Chicken (JFC)"
		request.OwnerType = "organization"
		request.Password = "123"

		err := authService.OwnerRegistration(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestOwnerRegistration_EmailAlreadyExist", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.OwnerRegistrationRequest{}
		request.OwnerEmail = "admin@jfc.com"
		request.OwnerName = "Jeremi Fried Chicken (JFC)"
		request.OwnerType = "organization"
		request.Password = "123"

		err := authService.OwnerRegistration(request)

		assert.Equal(t, "email already used", err.Error())
	})

	t.Run("TestOwnerLogin", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "admin@jfc.com"
		request.OwnerPassword = "123"

		_, err := authService.OwnerLogin(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestOwnerLogin_WrongPassword", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "admin@jfc.com"
		request.OwnerPassword = "1234"

		_, err := authService.OwnerLogin(request)

		assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())

	})

	t.Run("TestOwnerLogin_EmailNotExist", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "lost@jfc.com"
		request.OwnerPassword = "1234"

		_, err := authService.OwnerLogin(request)

		assert.Equal(t, "authentication error", err.Error())

	})

	t.Run("TestAgentLogin", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "budi@gmial.com"
		request.AgentPassword = "budi123"

		session, err := authService.AgentLogin(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		log.Println(session)
		assert.Equal(t, "agent-001", session.Id)
		assert.Equal(t, "agent", session.Role)
		assert.Equal(t, "owner-001", session.OwnerId)
		assert.Equal(t, "str-001", session.StoreId)

	})

	t.Run("TestAgentLogin_WrongPassword", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "budi@notexist.com"
		request.AgentPassword = "123"

		_, err := authService.AgentLogin(request)

		assert.Equal(t, "authentication error", err.Error())
	})

	t.Run("TestAgentLogin_EmailNotExist", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		joinRepo := &repository.JoinRepositoryImpl{}

		sessionList := make(map[string]*model.Session)
		agentSessionList := make(map[string]*model.AgentSession)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo,
			AgentRepo:        agentRepo,
			JoinRepo:         joinRepo,
			SessionList:      sessionList,
			AgentSessionList: agentSessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "jeremiah@smartveggiesmart.com"
		request.AgentPassword = "1234"

		_, err := authService.AgentLogin(request)

		assert.Equal(t, "authentication error", err.Error())
	})

	t.Run("CleanUpOwnerData", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}

		owner, _ := ownerRepo.FindByOwnerEmail("admin@jfc.com")

		_ = ownerRepo.Delete(owner.OwnerId)

	})

}
