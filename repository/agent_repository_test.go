package repository

import (
	"alingan/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentRepository(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	t.Run("TestFindAgentById", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		agent, err := agentRepository.FindAgentById("agent-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "agent-001", agent.AgentId)
		assert.Equal(t, "Budi", agent.AgentName)

	})

	t.Run("TestFindAgentByEmail", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		agent, err := agentRepository.FindAgentByEmail("budi@gmial.com")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "agent-001", agent.AgentId)
		assert.Equal(t, "Budi", agent.AgentName)

	})

	t.Run("TestInsert", func(t *testing.T) {

		data := entity.Agent{}
		data.AgentId = "agent-test"
		data.AgentName = "Agent Test"
		data.StoreId = "str-001"
		data.AgentEmail = "agent@test.com"
		data.AgentPassword = "test"

		agentRepository := &AgentRepositoryImpl{}

		err := agentRepository.Insert(data)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepository.FindAgentById("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "agent-test", agent.AgentId)
		assert.Equal(t, "Agent Test", agent.AgentName)
		assert.Equal(t, "agent@test.com", agent.AgentEmail)
		assert.Equal(t, "str-001", agent.StoreId)
		assert.Equal(t, "test", agent.AgentPassword)
	})

	t.Run("TestUpdate", func(t *testing.T) {

		data := entity.Agent{}
		data.AgentName = "Update Agent Test"
		data.AgentEmail = "updateagent@test.com"

		agentRepository := &AgentRepositoryImpl{}

		err := agentRepository.Update(data, "agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepository.FindAgentById("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "agent-test", agent.AgentId)
		assert.Equal(t, "Update Agent Test", agent.AgentName)
		assert.Equal(t, "updateagent@test.com", agent.AgentEmail)

	})

	t.Run("TestSetInactive", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		err := agentRepository.SetInactive("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepository.FindAgentById("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, agent.IsActive)

	})

	t.Run("TestSetActive", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		err := agentRepository.SetActive("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		agent, err := agentRepository.FindAgentById("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, agent.IsActive)

	})

	t.Run("TestCheckExist", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		result, err := agentRepository.CheckExist("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestCheckEmailExist", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		result, err := agentRepository.CheckEmailExist("updateagent@test.com")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestDelete", func(t *testing.T) {

		agentRepository := &AgentRepositoryImpl{}

		err := agentRepository.Delete("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		result, err := agentRepository.CheckExist("agent-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, result)

	})
}
