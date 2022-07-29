package service

import (
	"alingan/entity"
	"alingan/model"
	"alingan/repository"
	"alingan/util"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionService(t *testing.T) {

	ownerId := util.GenerateId("OWN")
	storeId := util.GenerateId("STR")
	agentId := util.GenerateId("AGT")
	productId1 := util.GenerateId("PRD")
	productId2 := util.GenerateId("PRD")

	storeRepo := &repository.StoreRepositoryImpl{}
	ownerRepo := &repository.OwnerRepositoryImpl{}
	productRepo := &repository.ProductRepositoryImpl{}
	transactionRepo := &repository.TransactionRepositoryImpl{}
	transactionItemRepo := &repository.TransactionItemRepositoryImpl{}
	agentRepo := &repository.AgentRepositoryImpl{}
	joinRepo := &repository.JoinRepositoryImpl{}
	testRepo := &repository.TestingRepository{}

	transactionService := &TransactionServiceImpl{
		StoreRepo:           storeRepo,
		OwnerRepo:           ownerRepo,
		ProductRepo:         productRepo,
		TransactionRepo:     transactionRepo,
		TransactionItemRepo: transactionItemRepo,
		AgentRepo:           agentRepo,
		JoinRepo:            joinRepo,
	}

	t.Run("PrepareOwnerData", func(t *testing.T) {

		owner := entity.Owner{}
		owner.OwnerId = ownerId
		owner.OwnerEmail = "jeremifriedchicken@jfc.com"
		owner.OwnerName = "Jeremi Fried Chicken"
		owner.OwnerType = "organization"
		owner.Password = "123"

		err := ownerRepo.Insert(owner)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
		}
	})

	t.Run("PrepareStoreData", func(t *testing.T) {

		store := entity.Store{}

		store.StoreId = storeId
		store.OwnerId = ownerId
		store.StoreName = "JFC - Tanggerang"
		store.StoreAddress = "Tanggerang"

		err := storeRepo.Insert(store)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
		}

	})

	t.Run("PrepareProductData1", func(t *testing.T) {

		product := entity.Product{}
		product.OwnerId = ownerId
		product.ProductId = productId1
		product.ProductName = "Chicken Breast Crispy"
		product.ProductMeasurementUnit = "pcs"
		product.ProductPrice = float64(15000)
		product.ProductStock = int64(40)

		err := productRepo.Insert(product)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("PrepareProductData2", func(t *testing.T) {

		product := entity.Product{}
		product.OwnerId = ownerId
		product.ProductId = productId2
		product.ProductName = "Wings Crispy"
		product.ProductMeasurementUnit = "pcs"
		product.ProductPrice = float64(12000)
		product.ProductStock = int64(40)

		err := productRepo.Insert(product)

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

	t.Run("TestCreateTransaction", func(t *testing.T) {

		item1 := model.CreateTransactionItemRequest{}
		item2 := model.CreateTransactionItemRequest{}

		item1.ProductId = productId1
		item1.UsedPrice = float64(15000)
		item1.BuyQuantity = int64(2)

		item2.ProductId = productId2
		item2.UsedPrice = float64(12000)
		item2.BuyQuantity = int64(2)

		request := model.CreateTransactionRequest{}
		request.StoreId = storeId
		request.AgentId = agentId
		request.TransactionDate = time.Now()

		request.Items = append(request.Items, item1)
		request.Items = append(request.Items, item2)

		err := transactionService.CreateTransaction(request)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestFindTransactionByOwner", func(t *testing.T) {

		transactions, err := transactionService.FindTransactionByOwner(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, storeId, transactions[0].StoreId)
		assert.Equal(t, agentId, transactions[0].AgentId)
		assert.Equal(t, float64(54000), transactions[0].TransactionTotal)

	})

	t.Run("CleanUpTransactionItemData", func(t *testing.T) {

		err := testRepo.DeleteAllTransactionItemByStore(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("CleanUpTransactionData", func(t *testing.T) {

		err := testRepo.DeleteAllTransactionByStore(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
	})

	t.Run("CleanUpProductData", func(t *testing.T) {

		err := testRepo.DeleteAllProductByOwner(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}
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
