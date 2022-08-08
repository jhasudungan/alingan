package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinRepository(t *testing.T) {

	t.Run("TestFindTransactionByOwnerId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindTransactionByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "TRX2b40ee07-f8dd-4b80-9956-e65b3a364443", results[0].TransactionId)
		assert.Equal(t, "agent-001", results[0].AgentId)
		assert.Equal(t, "str-001", results[0].StoreId)
	})

	t.Run("TestFindTransactionAgentAndStoreByTransactionId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		results, err := joinRepository.FindTransactionAgentAndStoreByTransactionId("TRX0a63cfe6-2761-4c12-a2c0-42051edcde10")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "TRX0a63cfe6-2761-4c12-a2c0-42051edcde10", results.TransactionId)
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
		assert.Equal(t, "Kapal Api ", results[0].ProductName)
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
		assert.Equal(t, "Jeremiah H.S", results[0].AgentName)
		assert.Equal(t, "Store 1 Smart Veggies", results[0].StoreName)

	})

	t.Run("TestFindOwnerByAgentId", func(t *testing.T) {

		joinRepository := &JoinRepositoryImpl{}

		result, err := joinRepository.FindOwnerByAgentId("AGT454497b9-74d1-4bb0-8753-962a962e31f6")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "owner-001", result.OwnerId)
		assert.Equal(t, "admin@smartveggiesmart.com", result.OwnerEmail)
		assert.Equal(t, "Smart Veggies Mart", result.OwnerName)

	})
}
