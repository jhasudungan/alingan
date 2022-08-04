package entity

import "time"

type Product struct {
	ProductId              string
	OwnerId                string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	IsActive               bool
	CreatedDate            time.Time
	LastModified           time.Time
}
