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

func TestProductService(t *testing.T) {

	productId := util.GenerateId("PRD")
	ownerId := util.GenerateId("OWN")

	testingRepo := &repository.TestingRepository{}
	ownerRepo := &repository.OwnerRepositoryImpl{}
	productRepo := &repository.ProductRepositoryImpl{}

	// service under test
	productService := &ProductServiceImpl{OwnerRepo: ownerRepo, ProductRepo: productRepo}

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

	t.Run("PrepareProductData", func(t *testing.T) {

		product := entity.Product{}
		product.OwnerId = ownerId
		product.ProductId = productId
		product.ProductName = "Chicken Breast Crispy"
		product.ProductMeasurementUnit = "pcs"
		product.ProductPrice = float64(150000)
		product.ProductStock = int64(40)

		err := productRepo.Insert(product)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

	})

	t.Run("TestCreateProductService", func(t *testing.T) {

		request := model.CreateProductRequest{}
		request.OwnerId = ownerId
		request.ProductName = "Wings Crispy Yakiniku"
		request.ProductPrice = float64(13000)
		request.ProductStock = int64(130)

		err := productService.CreateProduct(request)

		assert.Nil(t, err)

	})

	t.Run("TestUpdateProductService", func(t *testing.T) {

		request := model.UpdateProductRequest{}
		request.ProductName = "Chicken Breast Crispy Box"
		request.ProductMeasurementUnit = "box"
		request.ProductPrice = float64(25000)
		request.ProductStock = int64(60)

		err := productService.UpdateProduct(request, productId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			log.Print("After Update")
			t.FailNow()
		}

		product, err := productRepo.FindProductById(productId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			log.Print("After Query Test")
			t.FailNow()
		}

		assert.Equal(t, request.ProductName, product.ProductName)
		assert.Equal(t, request.ProductMeasurementUnit, product.ProductMeasurementUnit)
		assert.Equal(t, request.ProductPrice, product.ProductPrice)
		assert.Equal(t, request.ProductStock, product.ProductStock)
	})

	t.Run("FindProductById", func(t *testing.T) {

		product, err := productService.FindProductById(productId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "Chicken Breast Crispy Box", product.ProductName)
		assert.Equal(t, "box", product.ProductMeasurementUnit)
		assert.Equal(t, float64(25000), product.ProductPrice)
		assert.Equal(t, int64(60), product.ProductStock)
	})

	t.Run("TestFindProductByOwner", func(t *testing.T) {

		products, err := productService.FindProductByOwnerId(ownerId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, "Chicken Breast Crispy Box", products[0].ProductName)
		assert.Equal(t, "box", products[0].ProductMeasurementUnit)
		assert.Equal(t, float64(25000), products[0].ProductPrice)
		assert.Equal(t, int64(60), products[0].ProductStock)
	})

	t.Run("TesSetInactiveService", func(t *testing.T) {

		err := productService.SetProductInactive(productId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		product, err := productRepo.FindProductById(productId)

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		assert.Equal(t, false, product.IsActive)

	})

	t.Run("CleanUpStoreData", func(t *testing.T) {

		err := testingRepo.DeleteAllProductByOwner(ownerId)

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
