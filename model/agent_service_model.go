package model

type RegisterNewAgentRequest struct {
	StoreId       string
	AgentName     string
	AgentEmail    string
	AgentPassword string
}

type GetAgentInformationResponse struct {
	AgentId       string
	StoreId       string
	AgentName     string
	AgentEmail    string
	AgentPassword string
	IsActive      bool
	CreatedDate   string
	LastModified  string
}

type GetOwnerAgentListResponse struct {
	AgentId    string
	AgentName  string
	AgentEmail string
	StoreId    string
	StoreName  string
	IsActive   bool
}
