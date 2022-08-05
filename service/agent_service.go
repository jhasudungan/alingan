package service

import (
	"alingan/entity"
	"alingan/model"
	"alingan/repository"
	"alingan/util"
	"errors"
)

type AgentService interface {
	RegisterNewAgent(request model.RegisterNewAgentRequest) error
	GetAgentInformation(agentId string) (model.GetAgentInformationResponse, error)
	GetOwnerAgentList(ownerId string) ([]model.GetOwnerAgentListResponse, error)
	SetAgentInactive(agentId string) error
	SetAgentActive(agentId string) error
}

type AgentServiceImpl struct {
	AgentRepo repository.AgentRepository
	OwnerRepo repository.OwnerRepository
	JoinRepo  repository.JoinRepository
}

func (a *AgentServiceImpl) RegisterNewAgent(request model.RegisterNewAgentRequest) error {

	checkEmail, err := a.AgentRepo.CheckEmailExist(request.AgentEmail)

	if err != nil {
		return err
	}

	if checkEmail == true {
		return errors.New("email already used")
	}

	agent := entity.Agent{}
	agent.AgentId = util.GenerateId("AGT")
	agent.AgentEmail = request.AgentEmail
	agent.AgentName = request.AgentName
	agent.AgentPassword = request.AgentPassword
	agent.StoreId = request.StoreId

	err = a.AgentRepo.Insert(agent)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentServiceImpl) GetAgentInformation(agentId string) (model.GetAgentInformationResponse, error) {

	result := model.GetAgentInformationResponse{}

	agent, err := a.AgentRepo.FindAgentById(agentId)

	if err != nil {
		return result, err
	}

	result.AgentId = agent.AgentId
	result.StoreId = agent.StoreId
	result.AgentName = agent.AgentName
	result.AgentEmail = agent.AgentEmail
	result.AgentPassword = agent.AgentPassword
	result.IsActive = agent.IsActive
	result.CreatedDate = agent.CreatedDate
	result.LastModified = agent.LastModified

	return result, nil
}

func (a *AgentServiceImpl) SetAgentInactive(agentId string) error {

	checkExist, err := a.AgentRepo.CheckExist(agentId)

	if err != nil {
		return err
	}

	if checkExist == false {
		return errors.New("agent not exist")
	}

	err = a.AgentRepo.SetInactive(agentId)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentServiceImpl) SetAgentActive(agentId string) error {

	checkExist, err := a.AgentRepo.CheckExist(agentId)

	if err != nil {
		return err
	}

	if checkExist == false {
		return errors.New("agent not exist")
	}

	err = a.AgentRepo.SetActive(agentId)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentServiceImpl) GetOwnerAgentList(ownerId string) ([]model.GetOwnerAgentListResponse, error) {

	results := make([]model.GetOwnerAgentListResponse, 0)

	checkOwner, err := a.OwnerRepo.CheckExist(ownerId)

	if err != nil {
		return results, err
	}

	if checkOwner == false {
		return results, errors.New("owner is not exist")
	}

	agents, err := a.JoinRepo.FindAgentByOwnerId(ownerId)

	if err != nil {
		return results, err
	}

	for _, agent := range agents {

		data := model.GetOwnerAgentListResponse{}
		data.AgentId = agent.AgentId
		data.AgentName = agent.AgentName
		data.StoreId = agent.StoreId
		data.StoreName = agent.StoreName
		data.AgentEmail = agent.AgentEmail
		data.IsActive = agent.IsActive

		results = append(results, data)
	}

	return results, nil
}
