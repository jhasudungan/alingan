package entity

import "time"

type ProductImage struct {
	ProductImageId string
	ProductId      string
	LocationPath   string
	CreatedDate    time.Time
	LastModified   time.Time
}
