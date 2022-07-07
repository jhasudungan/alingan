package repository

import (
	"alingan/core/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnerRepository(t *testing.T) {

	t.Run("TestFindById", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		owner, err := ownerRepository.FindById("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "owner-001", owner.OwnerId)
		assert.Equal(t, "Smart Veggies Mart", owner.OwnerName)
		assert.Equal(t, "admin@smartveggiesmart.com", owner.OwnerEmail)
		assert.Equal(t, "organization", owner.OwnerType)
		assert.Equal(t, "smartveggies", owner.Password)

	})

	t.Run("TestInsert", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		data := entity.Owner{}
		data.OwnerId = "owner-test"
		data.OwnerName = "Owner Test"
		data.OwnerEmail = "ownet-test@mail.com"
		data.OwnerType = "organization"
		data.Password = "pwd123"

		err := ownerRepository.Insert(data)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		owner, err := ownerRepository.FindById("owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "owner-test", owner.OwnerId)
		assert.Equal(t, "Owner Test", owner.OwnerName)
		assert.Equal(t, "ownet-test@mail.com", owner.OwnerEmail)
		assert.Equal(t, "organization", owner.OwnerType)
		assert.Equal(t, "pwd123", owner.Password)

	})

	t.Run("TestUpdate", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		data := entity.Owner{}
		data.OwnerId = "owner-test"
		data.OwnerName = "Update Owner Test"
		data.OwnerType = "personal"

		err := ownerRepository.Update(data, "owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		owner, err := ownerRepository.FindById("owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "owner-test", owner.OwnerId)
		assert.Equal(t, "Update Owner Test", owner.OwnerName)
		assert.Equal(t, "personal", owner.OwnerType)
	})

	t.Run("TestCheckExist", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		result, err := ownerRepository.CheckExist("owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestCheckEmailExist", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		result, err := ownerRepository.CheckEmailExist("ownet-test@mail.com")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestDelete", func(t *testing.T) {

		ownerRepository := &OwnerRepositoryImpl{}

		err := ownerRepository.Delete("owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		result, err := ownerRepository.CheckExist("owner-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, result)

	})

}
