package model

import "time"

type FindTransactionByOwnerIdDTO struct {
	TransactionId    string
	TransactionDate  time.Time
	StoreId          string
	StoreName        string
	AgentId          string
	AgentName        string
	TransactionTotal float64
}

type FindAgentByOwnerIdDTO struct {
	AgentId    string
	AgentName  string
	AgentEmail string
	StoreId    string
	StoreName  string
	IsActive   bool
}

type FindTransactionAgentAndStoreByTransactionIdDTO struct {
	TransactionId    string
	TransactionDate  time.Time
	AgentId          string
	AgentName        string
	StoreId          string
	StoreName        string
	TransactionTotal string
}

type FindTransactionItemAndProductByTransactionIdDTO struct {
	TransactionItemId string
	TransactionId     string
	ProductId         string
	ProductName       string
	UsedPrice         float64
	BuyQuantity       int64
}
