package repository

import (
	"alingan/entity"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	t.Run("TestFindById", func(t *testing.T) {

		transactionRepository := &TransactionRepositoryImpl{}

		transaction, err := transactionRepository.FindById("trx-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-001", transaction.TransactionId)
		assert.Equal(t, "agent-001", transaction.AgentId)
		assert.Equal(t, "str-001", transaction.StoreId)
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
		data.TransactionTotal = float64(45000)

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

	t.Run("FindTransactionItemByTransactionId", func(t *testing.T) {

		transactionItemRepository := &TransactionItemRepositoryImpl{}

		items, err := transactionItemRepository.FindByTransactionId("trx-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "trx-item-001", items[0].TransactionItemId)
		assert.Equal(t, "prd-001", items[0].ProductId)
		assert.Equal(t, int64(2), items[0].BuyQuantity)
		assert.Equal(t, float64(2500), items[0].UsedPrice)

	})

	t.Run("TestInsertTransactionItem", func(t *testing.T) {

		transactionItemRepository := &TransactionItemRepositoryImpl{}

		item := entity.TransactionItem{}

		item.TransactionItemId = "trx-item-test"
		item.TransactionId = "trx-test"
		item.ProductId = "prd-001"
		item.BuyQuantity = int64(3)
		item.UsedPrice = float64(15000)

		err := transactionItemRepository.Insert(item)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		items, err := transactionItemRepository.FindByTransactionId("trx-test")

		assert.Equal(t, "trx-item-test", items[0].TransactionItemId)
		assert.Equal(t, "prd-001", items[0].ProductId)
		assert.Equal(t, int64(3), items[0].BuyQuantity)
		assert.Equal(t, float64(15000), items[0].UsedPrice)

	})

	t.Run("CleanUpTransaction", func(t *testing.T) {

		testingRepository := &TestingRepository{}

		testingRepository.DeleteTransactionById("trx-test")
		testingRepository.DeleteTransactionItemByTransactionId("trx-test")

	})
}
