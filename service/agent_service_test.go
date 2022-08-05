package service

import (
	"alingan/entity"
	"alingan/model"
	"alingan/repository"
	"alingan/util"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentService(t *testing.T) {

	ownerId := util.GenerateId("OWN")
	storeId := util.GenerateId("STR")
	agentId := util.GenerateId("AGT")

	storeRepo := &repository.StoreRepositoryImpl{}
	agentRepo := &repository.AgentRepositoryImpl{}
	ownerRepo := &repository.OwnerRepositoryImpl{}
	joinRepo := &repository.JoinRepositoryImpl{}
	testRepo := &repository.TestingRepository{}

	agentService := &AgentServiceImpl{
		OwnerRepo: ownerRepo,
		AgentRepo: agentRepo,
		JoinRepo:  joinRepo,
	}

	t.Run("PrepareOwnerData", func(t *testing.T) {

		owner := entity.Owner{}
		owner.OwnerId = ownerId
		owner.OwnerName = "Jeremi Fried Chicken (JFC)"
		owner.OwnerEmail = "admin@jfc.com"
		owner.OwnerType = "organization"
		owner.Password = "jfc123"

		err := ownerRepo.Insert(owner)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
	})

	t.Run("PrepareStoreDate", func(t *testing.T) {

		store := entity.Store{}
		store.StoreId = storeId
		store.OwnerId = ownerId
		store.StoreName = "JFC - Bogor"
		store.StoreAddress = "Bogor, Indonesia"

		err := storeRepo.Insert(store)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
	})

	t.Run("PrepareAgentData", func(t *testing.T) {

		agent := entity.Agent{}
		agent.AgentId = agentId
		agent.AgentName = "Jeremiah Hasudungan"
		agent.AgentPassword = "pwd123"
		agent.StoreId = storeId
		agent.AgentEmail = "jhasudungan@jfc.com"

		err := agentRepo.Insert(agent)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
	})

	t.Run("TestRegisterNewAgent", func(t *testing.T) {

		request := model.RegisterNewAgentRequest{}
		request.AgentName = "Elroi"
		request.StoreId = storeId
		request.AgentPassword = "123"
		request.AgentEmail = "elroi@jfc.com"

		err := agentService.RegisterNewAgent(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestCreateAgent_EmailIsAlreadyExist", func(t *testing.T) {

		request := model.RegisterNewAgentRequest{}
		request.AgentName = "Elroi"
		request.StoreId = "str-001"
		request.AgentPassword = "123"
		request.AgentEmail = "elroi@jfc.com"

		err := agentService.RegisterNewAgent(request)

		assert.Equal(t, "email already used", err.Error())
	})

	t.Run("TestSetAgentInactive", func(t *testing.T) {

		err := agentService.SetAgentInactive(agentId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepo.FindAgentById(agentId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, agent.IsActive)

	})

	t.Run("TestSetAgentActive", func(t *testing.T) {

		err := agentService.SetAgentActive(agentId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepo.FindAgentById(agentId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, agent.IsActive)

	})

	t.Run("TestGetOwnerAgentList", func(t *testing.T) {

		agents, err := agentService.GetOwnerAgentList(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, 2, len(agents))
	})

	t.Run("TestGetAgentInformation", func(t *testing.T) {

		agent, err := agentService.GetAgentInformation(agentId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, storeId, agent.StoreId)
		assert.Equal(t, "Jeremiah Hasudungan", agent.AgentName)
		assert.Equal(t, "jhasudungan@jfc.com", agent.AgentEmail)
	})

	t.Run("CleanUpStoreData", func(t *testing.T) {

		err := testRepo.DeleteAllStoreByOwner(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
	})

	t.Run("CleanUpAgentData", func(t *testing.T) {

		err := testRepo.DeleteAllAgentByStore(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("CleanUpOwnerData", func(t *testing.T) {

		err := ownerRepo.Delete(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

}
