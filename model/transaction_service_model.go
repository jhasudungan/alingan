package model

import "time"

type CreateTransactionRequest struct {
	TransactionDate time.Time
	AgentId         string
	StoreId         string
	Items           []CreateTransactionItemRequest
}

type CreateTransactionItemRequest struct {
	ProductId   string
	UsedPrice   float64
	BuyQuantity int64
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
