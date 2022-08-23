package repository

import (
	"alingan/entity"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductImageRepository(t *testing.T) {

	t.Run("TestInsert", func(t *testing.T) {

		productImageRepository := &ProductImageRepositoryImpl{}

		productImage := entity.ProductImage{}
		productImage.ProductImageId = "prd-image-test"
		productImage.ProductId = "prd-001"
		productImage.LocationPath = "https://via.placeholder.com/300"

		err := productImageRepository.Insert(productImage)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestFindProductImageById", func(t *testing.T) {

		productImageRepository := &ProductImageRepositoryImpl{}

		result, err := productImageRepository.FindProductImageById("prd-image-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-image-test", result.ProductImageId)
		assert.Equal(t, "prd-001", result.ProductId)
		assert.Equal(t, "https://via.placeholder.com/300", result.LocationPath)
	})

	t.Run("TestFindProductImageByProductId", func(t *testing.T) {

		productImageRepository := &ProductImageRepositoryImpl{}

		results, err := productImageRepository.FindProductImageByProductId("prd-001")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "prd-001", results[0].ProductId)
		assert.Equal(t, "prd-image-test", results[0].ProductImageId)
		assert.Equal(t, "https://via.placeholder.com/300", results[0].LocationPath)
	})

	t.Run("TestDelete", func(t *testing.T) {

		productImageRepository := &ProductImageRepositoryImpl{}

		err := productImageRepository.Delete("prd-image-test")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		log.Println("Check The DB")
	})

}
