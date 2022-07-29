package repository

import (
	"alingan/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreRepostiroy(t *testing.T) {

	t.Run("TestFindStoresByOwnerId", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		stores, err := storeRepository.FindStoresByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "str-002", stores[0].StoreId)
		assert.Equal(t, "Store 2 Smart Veggies", stores[0].StoreName)
		assert.Equal(t, true, stores[0].IsActive)
	})

	t.Run("TestFindStoreById", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		store, err := storeRepository.FindStoreById("str-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "str-001", store.StoreId)
		assert.Equal(t, "Store 1 Smart Veggies", store.StoreName)

	})

	t.Run("TestInsert", func(t *testing.T) {

		data := entity.Store{}
		data.StoreId = "str-test"
		data.StoreName = "Test Store"
		data.StoreAddress = "Test Store Address"

		storeRepository := &StoreRepositoryImpl{}

		err := storeRepository.Insert(data)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		store, err := storeRepository.FindStoreById("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "str-test", store.StoreId)
		assert.Equal(t, "Test Store", store.StoreName)
		assert.Equal(t, "Test Store Address", store.StoreAddress)
		assert.Equal(t, true, store.IsActive)
	})

	t.Run("TestUpdate", func(t *testing.T) {

		data := entity.Store{}
		data.StoreId = "str-test"
		data.StoreName = "Test Store Update"
		data.StoreAddress = "Test Store Update Address"

		storeRepository := &StoreRepositoryImpl{}

		err := storeRepository.Update(data, "str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		store, err := storeRepository.FindStoreById("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "str-test", store.StoreId)
		assert.Equal(t, "Test Store Update", store.StoreName)
		assert.Equal(t, "Test Store Update Address", store.StoreAddress)
		assert.Equal(t, true, store.IsActive)
	})

	t.Run("TestSetInactive", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		err := storeRepository.SetInactive("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		store, err := storeRepository.FindStoreById("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, store.IsActive)

	})

	t.Run("TestCheckExist", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		result, err := storeRepository.CheckExist("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)
	})

	t.Run("TestDelete", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		err := storeRepository.Delete("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		result, err := storeRepository.CheckExist("str-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, result)

	})

}
