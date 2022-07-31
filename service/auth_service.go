package service

import (
	"alingan/entity"
	"alingan/model"
	"alingan/repository"
	"alingan/util"
	"errors"
	"time"
)

type AuthService interface {
	OwnerRegistration(request model.OwnerRegistrationRequest) error
	OwnerLogin(request model.OwnerLoginRequest) (*model.Session, error)
	AgentLogin(request model.AgentLoginRequest) (*model.Session, error)
}

type AuthServiceImpl struct {
	SessionList map[string]*model.Session
	OwnerRepo   repository.OwnerRepository
	AgentRepo   repository.AgentRepository
}

func (a *AuthServiceImpl) OwnerRegistration(request model.OwnerRegistrationRequest) error {

	checkEmailExist, err := a.OwnerRepo.CheckEmailExist(request.OwnerEmail)

	if err != nil {
		return err
	}

	if checkEmailExist == true {
		return errors.New("email already used")
	}

	owner := entity.Owner{}
	owner.OwnerId = util.GenerateId("OWN")
	owner.OwnerEmail = request.OwnerEmail
	owner.OwnerName = request.OwnerName
	owner.OwnerType = request.OwnerType
	owner.Password = request.Password

	err = a.OwnerRepo.Insert(owner)

	if err != nil {
		return err
	}

	return nil
}

func (a *AuthServiceImpl) OwnerLogin(request model.OwnerLoginRequest) (*model.Session, error) {

	session := &model.Session{}

	checkEmailExist, err := a.OwnerRepo.CheckEmailExist(request.OwnerEmail)

	if err != nil {
		return session, err
	}

	if checkEmailExist == false {
		return session, errors.New("authentication error")
	}

	owner, err := a.OwnerRepo.FindByOwnerEmail(request.OwnerEmail)

	if err != nil {
		return session, err
	}

	if owner.Password != request.OwnerPassword {
		return session, errors.New("authentication error")
	}

	// session per login = 2 minutes
	session.Expiry = time.Now().Add(time.Minute * 2)
	session.Id = owner.OwnerId
	session.Role = "owner"
	session.Token = util.GenerateId("SES")

	// add session to session map with session.Id as key
	a.SessionList[session.Token] = session

	return session, nil
}

func (a *AuthServiceImpl) AgentLogin(request model.AgentLoginRequest) (*model.Session, error) {

	session := &model.Session{}

	checkEmailExist, err := a.AgentRepo.CheckEmailExist(request.AgentEmail)

	if err != nil {
		return session, err
	}

	if checkEmailExist == false {
		return session, errors.New("authentication error")
	}

	agent, err := a.AgentRepo.FindAgentByEmail(request.AgentEmail)

	if err != nil {
		return session, err
	}

	if agent.AgentPassword != request.AgentPassword {
		return session, errors.New("authentication error")
	}

	// session per login = 2 minutes
	session.Expiry = time.Now().Add(time.Minute * 2)
	session.Id = agent.AgentId
	session.Role = "agent"
	session.Token = util.GenerateId("SES")

	// add session to session map with session.Id as key
	a.SessionList[session.Token] = session

	return session, nil
}
