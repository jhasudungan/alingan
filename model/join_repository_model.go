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
