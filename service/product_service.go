package service

import (
	"alingan/core/entity"
	"alingan/core/model"
	"alingan/core/repository"
	"alingan/core/util"
	"errors"
)

type ProductService interface {
	CreateProduct(request model.CreateProductRequest) error
	UpdateProduct(request model.UpdateProductRequest, productId string) error
	FindProductByOwnerId(ownerId string) ([]model.FindProductByOwnerIdResponse, error)
	FindProductById(productId string) ([]model.FindProductByIdResponse, error)
	SetProductInactive(productId string) error
}

type ProductServiceImpl struct {
	ownerRepo   repository.OwnerRepository
	productRepo repository.ProductRepository
}

func (p *ProductServiceImpl) CreateProduct(request model.CreateProductRequest) error {

	product := entity.Product{}
	product.ProductId = util.GenerateId("PRD")
	product.OwnerId = request.OwnerId
	product.ProductName = request.ProductName
	product.ProductMeasurementUnit = request.ProductMeasurementUnit
	product.ProductPrice = request.ProductPrice
	product.ProductStock = request.ProductStock

	checkOwner, err := p.ownerRepo.CheckExist(request.OwnerId)

	if err != nil {
		return err
	}

	if checkOwner == false {
		return errors.New("owner is not exist")
	}

	err = p.productRepo.Insert(product)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductServiceImpl) UpdateProduct(request model.UpdateProductRequest, productId string) error {

	product := entity.Product{}
	product.ProductName = request.ProductName
	product.ProductMeasurementUnit = request.ProductMeasurementUnit
	product.ProductStock = request.ProductStock
	product.ProductPrice = request.ProductPrice

	checkProduct, err := p.productRepo.CheckExist(productId)

	if err != nil {
		return err
	}

	if checkProduct == false {
		return errors.New("product is not exist")
	}

	err = p.productRepo.Update(product, productId)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductServiceImpl) FindProductByOwnerId(ownerId string) ([]model.FindProductByOwnerIdResponse, error) {
	result := make([]model.FindProductByOwnerIdResponse, 0)

	checkOwner, err := p.ownerRepo.CheckExist(ownerId)

	if err != nil {
		return result, err
	}

	if checkOwner == false {
		return result, errors.New("owner is not exist")
	}

	products, err := p.productRepo.FindProductsByOwnerId(ownerId)

	if err != nil {
		return result, err
	}

	for _, product := range products {

		data := model.FindProductByOwnerIdResponse{}
		data.ProductId = product.ProductId
		data.OwnerId = product.OwnerId
		data.ProductName = product.ProductName
		data.ProductMeasurementUnit = product.ProductMeasurementUnit
		data.ProductPrice = product.ProductPrice
		data.ProductStock = product.ProductStock
		data.IsActive = product.IsActive
		data.LastModified = product.LastModified
		data.CreatedDate = product.CreatedDate

		result = append(result, data)
	}

	return result, nil
}

func (p *ProductServiceImpl) FindProductById(productId string) (model.FindProductByIdResponse, error) {
	result := model.FindProductByIdResponse{}

	checkProduct, err := p.productRepo.CheckExist(productId)

	if err != nil {
		return result, err
	}

	if checkProduct == false {
		return result, errors.New("product is not exist")
	}

	product, err := p.productRepo.FindProductById(productId)

	if err != nil {
		return result, err
	}

	result.ProductId = product.ProductId
	result.OwnerId = product.OwnerId
	result.ProductName = product.ProductName
	result.ProductMeasurementUnit = product.ProductMeasurementUnit
	result.ProductPrice = product.ProductPrice
	result.ProductStock = product.ProductStock
	result.IsActive = product.IsActive
	result.LastModified = product.LastModified
	result.CreatedDate = product.LastModified

	return result, nil
}

func (p *ProductServiceImpl) SetProductInactive(productId string) error {

	checkProduct, err := p.productRepo.CheckExist(productId)

	if err != nil {
		return err
	}

	if checkProduct == false {
		return errors.New("product is not exist")
	}

	err = p.productRepo.SetInactive(productId)

	if err != nil {
		return err
	}

	return nil
}