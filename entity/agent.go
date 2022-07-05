package entity

import "time"

type Agent struct {
	AgentId       string
	StoreId       string
	AgentName     string
	AgentEmail    string
	AgentPassword string
	IsActive      bool
	CreatedDate   time.Time
	LastModified  time.Time
}
