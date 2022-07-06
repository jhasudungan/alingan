package entity

import "time"

type Transaction struct {
	TransactionId    string
	TransactionDate  time.Time
	AgentId          string
	StoreId          string
	TransactionTotal float64
	CreatedDate      time.Time
	LastModified     time.Time
}
