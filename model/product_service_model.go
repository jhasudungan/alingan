package model

import "time"

type CreateProductRequest struct {
	OwnerId                string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	ProductStock           int64
}

type UpdateProductRequest struct {
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	ProductStock           int64
}

type FindProductByOwnerIdResponse struct {
	OwnerId                string
	ProductId              string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	ProductStock           int64
	IsActive               bool
	CreatedDate            time.Time
	LastModified           time.Time
	ImageUrl               string
}

type FindProductByIdResponse struct {
	OwnerId                string
	ProductId              string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	ProductStock           int64
	IsActive               bool
	CreatedDate            time.Time
	LastModified           time.Time
}
