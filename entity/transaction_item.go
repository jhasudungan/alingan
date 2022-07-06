package entity

import "time"

type TransactionItem struct {
	TransactionItemId string
	ProductId         string
	TransactionId     string
	UsedPrice         float64
	BuyQuantity       int64
	CreatedDate       time.Time
	LastModified      time.Time
}
