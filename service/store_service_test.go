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

func TestStoreService(t *testing.T) {

	storeId := util.GenerateId("STR")
	ownerId := util.GenerateId("OWN")

	testingRepo := &repository.TestingRepository{}
	storeRepo := &repository.StoreRepositoryImpl{}
	ownerRepo := &repository.OwnerRepositoryImpl{}

	// service under test
	storeService := &StoreServiceImpl{StoreRepo: storeRepo, OwnerRepo: ownerRepo}

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

	t.Run("PrepareStoreData", func(t *testing.T) {

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

	// Test The Service
	t.Run("TestCreateStoreService", func(t *testing.T) {

		request := model.CreateStoreRequest{}
		request.OwnerId = ownerId
		request.StoreName = "JFC - Jakarta"
		request.StoreAddress = "Jakarta, Indonesia"

		err := storeService.CreateStore(request)

		assert.Equal(t, nil, err)
	})

	t.Run("TestCreateStore_OwnerIsNotExist", func(t *testing.T) {

		request := model.CreateStoreRequest{}
		request.OwnerId = "OWN999"
		request.StoreName = "JFC - Jakarta"
		request.StoreAddress = "Jakarta, Indonesia"

		err := storeService.CreateStore(request)

		assert.Equal(t, "owner is not exist", err.Error())
	})

	t.Run("TestUpdateStore", func(t *testing.T) {

		request := model.UpdateStoreRequest{}
		request.StoreAddress = "Tanggerang, Indonesia"
		request.StoreName = "JFC - Tanggerang"

		err := storeService.UpdateStore(request, storeId)

		assert.Equal(t, nil, err)

		storeData, err := storeRepo.FindStoreById(storeId)

		assert.Equal(t, request.StoreName, storeData.StoreName)
		assert.Equal(t, request.StoreAddress, storeData.StoreAddress)

	})

	t.Run("TestUpdateStore_StoreIsNotExist", func(t *testing.T) {

		request := model.UpdateStoreRequest{}
		request.StoreAddress = "Tanggerang, Indonesia"
		request.StoreName = "JFC - Tanggerang"

		err := storeService.UpdateStore(request, "STR999")

		assert.Equal(t, "store is not exist", err.Error())

	})

	t.Run("FindStoreByOwnerId", func(t *testing.T) {

		stores, err := storeService.FindStoreByOwnerId(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, 2, len(stores))
		assert.Equal(t, ownerId, stores[1].OwnerId)
		assert.Equal(t, storeId, stores[1].StoreId)
		assert.Equal(t, "JFC - Tanggerang", stores[1].StoreName)
		assert.Equal(t, "Tanggerang, Indonesia", stores[1].StoreAddress)

	})

	t.Run("FindStoreByOwnerId_OwnerIsNotExist", func(t *testing.T) {

		_, err := storeService.FindStoreByOwnerId("OWN999")

		assert.Equal(t, "owner is not exist", err.Error())
	})

	t.Run("FindStoreById", func(t *testing.T) {

		store, err := storeService.FindStoreById(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, storeId, store.StoreId)
		assert.Equal(t, "JFC - Tanggerang", store.StoreName)
		assert.Equal(t, "Tanggerang, Indonesia", store.StoreAddress)

	})

	t.Run("FindStoreById", func(t *testing.T) {

		_, err := storeService.FindStoreById("STR999")

		assert.Equal(t, "store is not exist", err.Error())
	})

	t.Run("SetInactive", func(t *testing.T) {

		err := storeService.SetStoreInactive(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		store, err := storeService.FindStoreById(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, store.IsActive)
	})

	t.Run("SetActive", func(t *testing.T) {

		err := storeService.SetStoreActive(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		store, err := storeService.FindStoreById(storeId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, store.IsActive)
	})

	t.Run("CleanUpStoreData", func(t *testing.T) {

		err := testingRepo.DeleteAllStoreByOwner(ownerId)

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
