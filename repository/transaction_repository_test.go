package repository

import (
	"alingan/core/entity"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository(t *testing.T) {

	t.Run("TestFindById", func(t *testing.T) {

		transactionRepository := &TransactionRepositoryImpl{}

		transaction, err := transactionRepository.FindById("trx-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-001", transaction.TransactionId)

	})

	t.Run("TestInsert", func(t *testing.T) {

		transactionRepository := &TransactionRepositoryImpl{}

		layoutFormat := "2006-01-02 15:00:00"
		value := "2022-07-05 15:00:00"

		date, err := time.Parse(layoutFormat, value)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		data := entity.Transaction{}
		data.TransactionId = "trx-test"
		data.TransactionDate = date
		data.StoreId = "str-001"
		data.AgentId = "agent-001"
		data.TransactionTotal = float64(30000)

		err = transactionRepository.Insert(data)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		transaction, err := transactionRepository.FindById("trx-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-test", transaction.TransactionId)
	})
}
