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
	AgentLogin(request model.AgentLoginRequest) (*model.AgentSession, error)
	OwnerLogout(sessionToken string)
	AgentLogout(sessionToken string)
	GetOwnerProfileInformation(ownerId string) (model.GetOwnerProfileInformationResponse, error)
	UpdateOwnerProfile(request model.UpdateOwnerProfileRequest) error
}

type AuthServiceImpl struct {
	SessionList      map[string]*model.Session
	AgentSessionList map[string]*model.AgentSession
	JoinRepo         repository.JoinRepository
	OwnerRepo        repository.OwnerRepository
	AgentRepo        repository.AgentRepository
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
	session.Expiry = time.Now().Add(time.Minute * 30)
	session.Id = owner.OwnerId
	session.Role = "owner"
	session.Token = util.GenerateId("SES")

	// add session to session map with session.Id as key
	a.SessionList[session.Token] = session

	return session, nil
}

func (a *AuthServiceImpl) AgentLogin(request model.AgentLoginRequest) (*model.AgentSession, error) {

	session := &model.AgentSession{}

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

	owner, err := a.JoinRepo.FindOwnerByAgentId(agent.AgentId)

	if err != nil {
		return session, err
	}

	// session per login = 30 minutes
	session.Expiry = time.Now().Add(time.Minute * 30)
	session.Id = agent.AgentId
	session.Role = "agent"
	session.Token = util.GenerateId("SES")
	session.StoreId = agent.StoreId
	session.OwnerId = owner.OwnerId

	// add session to session map with session.Id as key
	a.AgentSessionList[session.Token] = session

	return session, nil
}

func (a *AuthServiceImpl) OwnerLogout(sessionToken string) {
	delete(a.SessionList, sessionToken)
}

func (a *AuthServiceImpl) AgentLogout(sessionToken string) {
	delete(a.AgentSessionList, sessionToken)
}

func (a *AuthServiceImpl) GetOwnerProfileInformation(ownerId string) (model.GetOwnerProfileInformationResponse, error) {

	result := model.GetOwnerProfileInformationResponse{}

	checkEmailExist, err := a.OwnerRepo.CheckExist(ownerId)

	if err != nil {
		return result, err
	}

	if checkEmailExist == false {
		return result, errors.New("owner is not exist")
	}

	ownerData, err := a.OwnerRepo.FindById(ownerId)

	result.OwnerId = ownerData.OwnerId
	result.OwnerEmail = ownerData.OwnerEmail
	result.OwnerName = ownerData.OwnerName
	result.Password = ownerData.Password
	result.IsActive = ownerData.IsActive
	result.CreatedDate = ownerData.CreatedDate.Format("2006-01-02 15:04:05")
	result.LastModified = ownerData.LastModified.Format("2006-01-02 15:04:05")
	result.OwnerType = ownerData.OwnerType

	return result, nil
}

func (a *AuthServiceImpl) UpdateOwnerProfile(request model.UpdateOwnerProfileRequest) error {

	checkExist, err := a.OwnerRepo.CheckExist(request.OwnerId)

	if err != nil {
		return err
	}

	if checkExist == false {
		return errors.New("owner is not exist")
	}

	owner := entity.Owner{}
	owner.OwnerId = request.OwnerId
	owner.OwnerName = request.OwnerName
	owner.Password = request.Password

	err = a.OwnerRepo.Update(owner, owner.OwnerId)

	if err != nil {
		return err
	}

	return nil
}
