package service

import (
	"alingan/core/model"
	"alingan/core/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {

	t.Run("TestOwnerLogin", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "admin@smartveggiesmart.com"
		request.OwnerPassword = "smartveggies"

		session, err := authService.OwnerLogin(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		log.Println(session)
		assert.Equal(t, "owner-001", session.Id)
		assert.Equal(t, "owner", session.Role)

	})

	t.Run("TestOwnerLogin_WrongPassword", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "admin@smartveggiesmart.com"
		request.OwnerPassword = "1234"

		_, err := authService.OwnerLogin(request)

		assert.Equal(t, "authentication error", err.Error())

	})

	t.Run("TestOwnerLogin_EmailNotExist", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.OwnerLoginRequest{}
		request.OwnerEmail = "lost@smartveggiesmart.com"
		request.OwnerPassword = "1234"

		_, err := authService.OwnerLogin(request)

		assert.Equal(t, "authentication error", err.Error())

	})

	t.Run("TestAgentLogin", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "jeremiahhs@smartveggiesmart.com"
		request.AgentPassword = "jhs123"

		session, err := authService.AgentLogin(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		log.Println(session)
		assert.Equal(t, "agent-001", session.Id)
		assert.Equal(t, "agent", session.Role)

	})

	t.Run("TestAgentLogin_WrongPassword", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "jeremiahhs@smartveggiesmart.com"
		request.AgentPassword = "1234"

		_, err := authService.AgentLogin(request)

		assert.Equal(t, "authentication error", err.Error())
	})

	t.Run("TestAgentLogin_EmailNotExist", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.AgentLoginRequest{}
		request.AgentEmail = "jeremiah@smartveggiesmart.com"
		request.AgentPassword = "1234"

		_, err := authService.AgentLogin(request)

		assert.Equal(t, "authentication error", err.Error())
	})

	t.Run("TestOwnerRegistration", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}
		agentRepo := &repository.AgentRepositoryImpl{}
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

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
		sessionList := make(map[string]*model.Session)

		authService := &AuthServiceImpl{OwnerRepo: ownerRepo, AgentRepo: agentRepo, SessionList: sessionList}

		request := model.OwnerRegistrationRequest{}
		request.OwnerEmail = "admin@jfc.com"
		request.OwnerName = "Jeremi Fried Chicken (JFC)"
		request.OwnerType = "organization"
		request.Password = "123"

		err := authService.OwnerRegistration(request)

		assert.Equal(t, "email already used", err.Error())
	})

	t.Run("CleanUpOwnerData", func(t *testing.T) {

		ownerRepo := &repository.OwnerRepositoryImpl{}

		owner, _ := ownerRepo.FindByOwnerEmail("admin@jfc.com")

		_ = ownerRepo.Delete(owner.OwnerId)

	})

}
