package service

import (
	"alingan/entity"
	"alingan/model"
	"alingan/repository"
	"alingan/util"
	"errors"
	"time"
)

type TransactionService interface {
	CreateTransaction(request model.CreateTransactionRequest) error
	CountTotalTransaction(request model.CreateTransactionRequest) float64
	FindTransactionByOwner(ownerId string) ([]model.FindTransactionByOwnerResponse, error)
	FindTransactionByAgent(agentId string) ([]model.FindTransactionByAgentResponse, error)
	FindTransactionByStore(ownerId string) ([]model.FindTransactionByStoreResponse, error)
	GetTransactionInformation(transactionId string) (model.GetTransactionInformationResponse, error)
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

	transaction.TransactionDate = time.Now()
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
		data.TransactionDate = transaction.TransactionDate.Format("2006-01-02 15:04:05")
		data.AgentId = transaction.AgentId
		data.AgentName = transaction.AgentName
		data.StoreName = transaction.StoreName
		data.StoreId = transaction.StoreId
		data.TransactionTotal = transaction.TransactionTotal

		results = append(results, data)
	}

	return results, nil
}

func (t *TransactionServiceImpl) FindTransactionByAgent(agentId string) ([]model.FindTransactionByAgentResponse, error) {

	results := make([]model.FindTransactionByAgentResponse, 0)

	checkOwner, err := t.AgentRepo.CheckExist(agentId)

	if err != nil {
		return results, nil
	}

	if checkOwner == false {
		return results, errors.New("agent is not exist")
	}

	transactions, err := t.JoinRepo.FindTransactionByAgentId(agentId)

	if err != nil {
		return results, err
	}

	for _, transaction := range transactions {

		data := model.FindTransactionByAgentResponse{}

		data.TransactionId = transaction.TransactionId
		data.TransactionDate = transaction.TransactionDate.Format("2006-01-02 15:04:05")
		data.AgentId = transaction.AgentId
		data.AgentName = transaction.AgentName
		data.StoreName = transaction.StoreName
		data.StoreId = transaction.StoreId
		data.TransactionTotal = transaction.TransactionTotal

		results = append(results, data)
	}

	return results, nil
}

func (t *TransactionServiceImpl) FindTransactionByStore(storeId string) ([]model.FindTransactionByStoreResponse, error) {

	results := make([]model.FindTransactionByStoreResponse, 0)

	checkOwner, err := t.StoreRepo.CheckExist(storeId)

	if err != nil {
		return results, nil
	}

	if checkOwner == false {
		return results, errors.New("agent is not exist")
	}

	transactions, err := t.JoinRepo.FindTransactionByStoreId(storeId)

	if err != nil {
		return results, err
	}

	for _, transaction := range transactions {

		data := model.FindTransactionByStoreResponse{}

		data.TransactionId = transaction.TransactionId
		data.TransactionDate = transaction.TransactionDate.Format("2006-01-02 15:04:05")
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

func (t *TransactionServiceImpl) GetTransactionInformation(transactionId string) (model.GetTransactionInformationResponse, error) {

	result := model.GetTransactionInformationResponse{}

	transaction, err := t.JoinRepo.FindTransactionAgentAndStoreByTransactionId(transactionId)

	if err != nil {
		return result, err
	}

	result.TransactionId = transaction.TransactionId
	result.StoreId = transaction.StoreId
	result.StoreName = transaction.StoreName
	result.AgentId = transaction.AgentId
	result.AgentName = transaction.AgentName
	result.TransactionDate = transaction.TransactionDate.Format("2006-01-02 15:04:05")

	transactionItems, err := t.JoinRepo.FindTransactionItemAndProductByTransactionId(transactionId)

	for _, transactionItem := range transactionItems {

		item := model.GetTransactionInformationTransactionItem{}

		item.ProductName = transactionItem.ProductName
		item.ProductId = transactionItem.ProductId
		item.BuyQuantity = transactionItem.BuyQuantity
		item.UsedPrice = transactionItem.UsedPrice

		result.Items = append(result.Items, item)
	}

	return result, nil
}
