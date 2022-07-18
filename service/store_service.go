package service

import (
	"alingan/core/entity"
	"alingan/core/model"
	"alingan/core/repository"
	"alingan/core/util"
	"errors"
)

type StoreService interface {
	CreateStore(request model.CreateStoreRequest) error
	UpdateStore(request model.UpdateStoreRequest, storeId string) error
	FindStoreByOwnerId(ownerId string) ([]model.FindStoreByOwnerIdResponse, error)
	FindStoreById(storeId string) (model.FindStoreByIdResponse, error)
	SetStoreInactive(storeId string) error
}

type StoreServiceImpl struct {
	OwnerRepo repository.OwnerRepository
	StoreRepo repository.StoreRepository
}

func (s *StoreServiceImpl) CreateStore(request model.CreateStoreRequest) error {

	id := util.GenerateId("STR")

	store := entity.Store{}

	store.StoreId = id
	store.OwnerId = request.OwnerId
	store.StoreName = request.StoreName
	store.StoreAddress = request.StoreAddress

	checkOwner, err := s.OwnerRepo.CheckExist(request.OwnerId)

	if err != nil {
		return err
	}

	if checkOwner == false {
		return errors.New("owner is not exist")
	}

	err = s.StoreRepo.Insert(store)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreServiceImpl) UpdateStore(request model.UpdateStoreRequest, storeId string) error {

	checkExist, err := s.StoreRepo.CheckExist(storeId)

	if err != nil {
		return err
	}

	if checkExist == false {
		return errors.New("store is not exist")
	}

	store := entity.Store{}
	store.StoreName = request.StoreName
	store.StoreAddress = request.StoreAddress

	err = s.StoreRepo.Update(store, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreServiceImpl) FindStoreByOwnerId(ownerId string) ([]model.FindStoreByOwnerIdResponse, error) {

	results := make([]model.FindStoreByOwnerIdResponse, 0)

	checkOwner, err := s.OwnerRepo.CheckExist(ownerId)

	if err != nil {
		return results, err
	}

	if checkOwner == false {
		return results, errors.New("owner is not exist")
	}

	stores, err := s.StoreRepo.FindStoresByOwnerId(ownerId)

	if err != nil {
		return results, err
	}

	for _, store := range stores {

		data := model.FindStoreByOwnerIdResponse{}
		data.OwnerId = store.OwnerId
		data.StoreId = store.StoreId
		data.StoreName = store.StoreName
		data.StoreAddress = store.StoreAddress
		data.IsActive = store.IsActive
		data.CreatedDate = store.CreatedDate
		data.LastModified = store.LastModified

		results = append(results, data)
	}

	return results, nil

}

func (s *StoreServiceImpl) FindStoreById(storeId string) (model.FindStoreByIdResponse, error) {

	result := model.FindStoreByIdResponse{}

	checkStore, err := s.StoreRepo.CheckExist(storeId)

	if err != nil {
		return result, err
	}

	if checkStore == false {
		return result, errors.New("store is not exist")
	}

	store, err := s.StoreRepo.FindStoreById(storeId)

	if err != nil {
		return result, err
	}

	result.StoreId = store.StoreId
	result.OwnerId = store.OwnerId
	result.StoreName = store.StoreName
	result.StoreAddress = store.StoreAddress
	result.IsActive = store.IsActive
	result.CreatedDate = store.CreatedDate
	result.LastModified = store.LastModified

	return result, nil
}

func (s *StoreServiceImpl) SetStoreInactive(storeId string) error {

	checkExist, err := s.StoreRepo.CheckExist(storeId)

	if err != nil {
		return err
	}

	if checkExist == false {
		return errors.New("store is not exist")
	}

	err = s.StoreRepo.SetInactive(storeId)

	if err != nil {
		return err
	}

	return nil
}
