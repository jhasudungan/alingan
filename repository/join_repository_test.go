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

		assert.Equal(t, "trx-001", results[0].TransactionId)
		assert.Equal(t, "agent-001", results[0].AgentId)
		assert.Equal(t, "str-001", results[0].StoreId)
	})
}
