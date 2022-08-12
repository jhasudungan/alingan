package model

type FindOwnerMostPurchasedProductByQuantityDTO struct {
	ProductId      string
	ProductName    string
	TotalPurchased int64
}

type FindOwnerMostPurchasedProductByRevenueDTO struct {
	ProductId    string
	ProductName  string
	TotalRevenue float64
}

type FindOwnerStoreWithTheMostTransactionDTO struct {
	StoreId               string
	StoreName             string
	StoreTotalTransaction int64
}

type FindOwnerAgentWithTheMostTransactionDTO struct {
	AgentId                 string
	AgentName               string
	TotalTransactionHandled int64
}
