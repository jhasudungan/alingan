package service

import (
	"alingan/core/entity"
	"alingan/core/model"
	"alingan/core/repository"
	"alingan/core/util"
	"errors"
)

type TransactionService interface {
	CreateTransaction(request model.CreateTransactionRequest) error
	CountTotalTransaction(request model.CreateTransactionRequest) float64
}

type TransactionServiceImpl struct {
	StoreRepo           repository.StoreRepository
	OwnerRepo           repository.OwnerRepository
	ProductRepo         repository.ProductRepository
	TransactionRepo     repository.TransactionRepository
	TransactionItemRepo repository.TransactionItemRepository
	AgentRepo           repository.AgentRepository
	JoinRepo            repository.JoinRepository
}

func (t *TransactionServiceImpl) CreateTransaction(request model.CreateTransactionRequest) error {

	checkStore, err := t.StoreRepo.CheckExist(request.StoreId)

	if err != nil {
		return err
	}

	if checkStore == false {
		return errors.New("store is not exist")
	}

	checkAgent, err := t.AgentRepo.CheckExist(request.AgentId)

	if err != nil {
		return err
	}

	if checkAgent == false {
		return errors.New("agent is not exist")
	}

	transactionId := util.GenerateId("TRX")
	transaction := entity.Transaction{}

	transaction.TransactionId = transactionId
	transaction.TransactionTotal = t.CountTotalTransaction(request)
	transaction.AgentId = request.AgentId
	transaction.StoreId = request.StoreId

	err = t.TransactionRepo.Insert(transaction)

	if err != nil {
		return err
	}

	for _, item := range request.Items {

		transactionItem := entity.TransactionItem{}

		transactionItem.TransactionId = transactionId
		transactionItem.TransactionItemId = util.GenerateId("TRX-ITEM")
		transactionItem.BuyQuantity = item.BuyQuantity
		transactionItem.UsedPrice = item.UsedPrice
		transactionItem.ProductId = item.ProductId

		err = t.TransactionItemRepo.Insert(transactionItem)

		if err != nil {
			return err
		}

	}

	return nil
}

func (t *TransactionServiceImpl) FindTransactionByOwner(ownerId string) ([]model.FindTransactionByOwnerResponse, error) {

	results := make([]model.FindTransactionByOwnerResponse, 0)

	checkOwner, err := t.OwnerRepo.CheckExist(ownerId)

	if err != nil {
		return results, nil
	}

	if checkOwner == false {
		return results, errors.New("owner is not exist")
	}

	transactions, err := t.JoinRepo.FindTransactionByOwnerId(ownerId)

	if err != nil {
		return results, err
	}

	for _, transaction := range transactions {

		data := model.FindTransactionByOwnerResponse{}

		data.TransactionId = transaction.TransactionId
		data.TransactionDate = transaction.TransactionDate
		data.AgentId = transaction.AgentId
		data.AgentName = transaction.AgentName
		data.StoreName = transaction.StoreName
		data.StoreId = transaction.StoreId
		data.TransactionTotal = transaction.TransactionTotal

		results = append(results, data)
	}

	return results, nil
}

func (t *TransactionServiceImpl) CountTotalTransaction(request model.CreateTransactionRequest) float64 {

	result := float64(0)

	for _, item := range request.Items {
		result = result + (item.UsedPrice * float64(item.BuyQuantity))
	}

	return result
}
