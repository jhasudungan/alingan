package repository

import (
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

		assert.Equal(t, "str-002", stores[0]["storeId"])
		assert.Equal(t, "Store 2 Smart Veggies", stores[0]["storeName"])
		assert.Equal(t, true, stores[0]["isActive"])
	})

	t.Run("TestFindStoreById", func(t *testing.T) {

		storeRepository := &StoreRepositoryImpl{}

		store, err := storeRepository.FindStoreById("str-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "str-001", store["storeId"])
		assert.Equal(t, "Store 1 Smart Veggies", store["storeName"])

	})

	t.Run("TestInsert", func(t *testing.T) {

		data := make(map[string]interface{})

		data["storeId"] = "str-test"
		data["ownerId"] = "owner-001"
		data["storeName"] = "Test Store"
		data["storeAddress"] = "Test Store Address"

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

		assert.Equal(t, "str-test", store["storeId"])
		assert.Equal(t, "Test Store", store["storeName"])
		assert.Equal(t, "Test Store Address", data["storeAddress"])
		assert.Equal(t, true, store["isActive"])
	})

	t.Run("TestUpdate", func(t *testing.T) {

		data := make(map[string]interface{})

		data["storeName"] = "Test Store Update"
		data["storeAddress"] = "Test Store Update Address"

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

		assert.Equal(t, "str-test", store["storeId"])
		assert.Equal(t, "Test Store Update", store["storeName"])
		assert.Equal(t, "Test Store Update Address", store["storeAddress"])
		assert.Equal(t, true, store["isActive"])
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

		assert.Equal(t, false, store["isActive"])

	})

}
