package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinRepository(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	t.Run("TestFindTransactionByOwnerId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindTransactionByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-002", results[0].TransactionId)
		assert.Equal(t, "agent-001", results[0].AgentId)
		assert.Equal(t, "str-001", results[0].StoreId)
	})

	t.Run("TestFindTransactionAgentAndStoreByTransactionId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindTransactionAgentAndStoreByTransactionId("trx-002")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-002", results.TransactionId)
		assert.Equal(t, "agent-001", results.AgentId)
		assert.Equal(t, "str-001", results.StoreId)
	})

	t.Run("TestFindTransactionItemAndProductByTransactionId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindTransactionItemAndProductByTransactionId("trx-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-001", results[0].TransactionId)
		assert.Equal(t, "prd-001", results[0].ProductId)
		assert.Equal(t, "Indomie Goreng", results[0].ProductName)
	})

	t.Run("TestFindAgentByOwnerId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindAgentByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "agent-001", results[0].AgentId)
		assert.Equal(t, "str-001", results[0].StoreId)
		assert.Equal(t, "Budi", results[0].AgentName)
		assert.Equal(t, "Toko Prima 1 - Salemba", results[0].StoreName)

	})

	t.Run("TestFindOwnerByAgentId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		result, err := joinRepository.FindOwnerByAgentId("agent-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "owner-001", result.OwnerId)
		assert.Equal(t, "tokoprima@gmail.co.id", result.OwnerEmail)
		assert.Equal(t, "Toko Prima", result.OwnerName)

	})
}
