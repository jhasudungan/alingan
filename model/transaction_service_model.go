package model

import "time"

type CreateTransactionRequest struct {
	TransactionDate time.Time
	AgentId         string                         `json:"agentId"`
	StoreId         string                         `json:"storeId"`
	Items           []CreateTransactionItemRequest `json:"transactionItems"`
}

type CreateTransactionItemRequest struct {
	ProductId   string  `json:"productId"`
	UsedPrice   float64 `json:"usedPrice"`
	BuyQuantity int64   `json:"buyQuantity"`
}

type FindTransactionByOwnerResponse struct {
	TransactionId    interface{}
	TransactionDate  interface{}
	AgentId          interface{}
	AgentName        interface{}
	StoreId          interface{}
	StoreName        interface{}
	TransactionTotal interface{}
}

type GetTransactionInformationResponse struct {
	TransactionId    string
	TransactionDate  string
	AgentId          string
	AgentName        string
	StoreId          string
	StoreName        string
	TransactionTotal string
	Items            []GetTransactionInformationTransactionItem
}

type GetTransactionInformationTransactionItem struct {
	ProductId   string
	ProductName string
	UsedPrice   float64
	BuyQuantity int64
}
