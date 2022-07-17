package model

import "time"

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
	CreatedDate   time.Time
	LastModified  time.Time
}

type GetOwnerAgentListResponse struct {
	AgentId    string
	AgentName  string
	AgentEmail string
	StoreId    string
	StoreName  string
	IsActive   bool
}
