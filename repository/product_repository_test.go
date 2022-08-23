package repository

import (
	"alingan/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	t.Run("TestFindProductsByOwnerId", func(t *testing.T) {

		productRepository := ProductRepositoryImpl{}

		products, err := productRepository.FindProductsByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-002", products[0].ProductId)
		assert.Equal(t, "Telur Ayam", products[0].ProductName)
	})

	t.Run("TestFindProductById", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		product, err := productRepository.FindProductById("prd-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-001", product.ProductId)
		assert.Equal(t, "Indomie Goreng", product.ProductName)

	})

	t.Run("TestInsert", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		data := entity.Product{}
		data.ProductId = "prd-test"
		data.OwnerId = "owner-001"
		data.ProductName = "Product Test"
		data.ProductMeasurementUnit = "pcs"
		data.ProductPrice = float64(15000)

		err := productRepository.Insert(data)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepository.FindProductById("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-test", product.ProductId)
		assert.Equal(t, "owner-001", product.OwnerId)
		assert.Equal(t, "Product Test", product.ProductName)
		assert.Equal(t, "pcs", product.ProductMeasurementUnit)
		assert.Equal(t, float64(15000), product.ProductPrice)

	})

	t.Run("TestUpdate", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		data := entity.Product{}
		data.ProductId = "prd-test"
		data.ProductName = "Product Update Test"
		data.ProductMeasurementUnit = "box"
		data.ProductPrice = float64(13000)

		err := productRepository.Update(data, data.ProductId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepository.FindProductById("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "Product Update Test", product.ProductName)
		assert.Equal(t, "box", product.ProductMeasurementUnit)
		assert.Equal(t, float64(13000), product.ProductPrice)

	})

	t.Run("TestSetInactive", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		err := productRepository.SetActive("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepository.FindProductById("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, product.IsActive)
	})

	t.Run("TestSetActive", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		err := productRepository.SetActive("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepository.FindProductById("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, product.IsActive)
	})

	t.Run("TestCheckExist", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		result, err := productRepository.CheckExist("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, true, result)

	})

	t.Run("TestDelete", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		err := productRepository.Delete("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		result, err := productRepository.CheckExist("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, result)

	})
}
