package model

import "time"

type CreateStoreRequest struct {
	OwnerId      string
	StoreName    string
	StoreAddress string
}

type UpdateStoreRequest struct {
	StoreName    string
	StoreAddress string
}

type FindStoreByIdResponse struct {
	StoreId      string
	OwnerId      string
	StoreName    string
	StoreAddress string
	IsActive     bool
	CreatedDate  time.Time
	LastModified time.Time
}

type FindStoreByOwnerIdResponse struct {
	StoreId      string
	OwnerId      string
	StoreName    string
	StoreAddress string
	IsActive     bool
	CreatedDate  time.Time
	LastModified time.Time
}
