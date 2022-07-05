package entity

import "time"

type Store struct {
	StoreId      string
	OwnerId      string
	StoreName    string
	StoreAddress string
	IsActive     bool
	CreatedDate  time.Time
	LastModified time.Time
}
