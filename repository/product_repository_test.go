package repository

import (
	"alingan/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {

	/**
	Make sure below data available in core.store
		- prd-001 (Kapal Api)
		- prd-002 (Torabika Creamy Latte)
		- PRDa543809e-f36f-443a-a815-64c0e2f0e09c (Abc Susu) and this one should be the last inserted
	*/

	t.Run("TestFindProductsByOwnerId", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		products, err := productRepository.FindProductsByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "PRDa543809e-f36f-443a-a815-64c0e2f0e09c", products[0].ProductId)
		assert.Equal(t, "Abc Susu ", products[0].ProductName)
	})

	t.Run("TestFindProductById", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		product, err := productRepository.FindProductById("prd-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-001", product.ProductId)
		assert.Equal(t, "Kapal Api ", product.ProductName)

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
