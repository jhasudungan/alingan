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
