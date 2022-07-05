package repository

import (
	"alingan/core/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {

	t.Run("TestFindProductsByOwnerId", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		products, err := productRepository.FindProductsByOwnerId("owner-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-001", products[0].ProductId)
		assert.Equal(t, "Fresh Orange Juice", products[0].ProductName)
	})

	t.Run("TestFindProductById", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		product, err := productRepository.FindProductById("prd-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-001", product.ProductId)
		assert.Equal(t, "Fresh Orange Juice", product.ProductName)

	})

	t.Run("TestInsert", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		data := entity.Product{}
		data.ProductId = "prd-test"
		data.OwnerId = "owner-001"
		data.ProductName = "Product Test"
		data.ProductMeasurementUnit = "pcs"
		data.ProductPrice = float64(15000)
		data.ProductStock = int64(100)

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
		assert.Equal(t, int64(100), product.ProductStock)

	})

	t.Run("TestUpdate", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		data := entity.Product{}
		data.ProductId = "prd-test"
		data.ProductName = "Product Update Test"
		data.ProductMeasurementUnit = "box"
		data.ProductPrice = float64(13000)
		data.ProductStock = int64(150)

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
		assert.Equal(t, int64(150), product.ProductStock)

	})

	t.Run("TestSetInactive", func(t *testing.T) {

		productRepository := &ProductRepositoryImpl{}

		err := productRepository.SetInactive("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepository.FindProductById("prd-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, product.IsActive)
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
