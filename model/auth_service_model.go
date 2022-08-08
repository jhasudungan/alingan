package model

import "time"

type Session struct {
	Id     string
	Role   string
	Expiry time.Time
	Token  string
}

type AgentSession struct {
	Id      string
	StoreId string
	OwnerId string
	Role    string
	Expiry  time.Time
	Token   string
}

type OwnerRegistrationRequest struct {
	OwnerEmail string
	Password   string
	OwnerName  string
	OwnerType  string
}

type OwnerLoginRequest struct {
	OwnerEmail    string
	OwnerPassword string
}
type AgentLoginRequest struct {
	AgentEmail    string
	AgentPassword string
}
