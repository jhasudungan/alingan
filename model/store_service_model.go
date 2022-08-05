package model

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
	CreatedDate  string
	LastModified string
}

type FindStoreByOwnerIdResponse struct {
	StoreId      string
	OwnerId      string
	StoreName    string
	StoreAddress string
	IsActive     bool
	CreatedDate  string
	LastModified string
}
