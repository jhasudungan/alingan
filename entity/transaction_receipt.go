package entity

import "time"

type TransactionReceipt struct {
	TransactionReceiptId string
	TransactionId        string
	LocationPath         string
	CreatedDate          time.Time
	LastModified         time.Time
}
