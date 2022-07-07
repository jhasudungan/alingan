package entity

import "time"

type Owner struct {
	OwnerId      string
	OwnerName    string
	OwnerType    string
	OwnerEmail   string
	Password     string
	IsActive     bool
	CreatedDate  time.Time
	LastModified time.Time
}
