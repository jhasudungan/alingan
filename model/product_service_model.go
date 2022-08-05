package model

type CreateProductRequest struct {
	OwnerId                string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
}

type UpdateProductRequest struct {
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
}

type FindProductByOwnerIdResponse struct {
	OwnerId                string
	ProductId              string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	IsActive               bool
	CreatedDate            string
	LastModified           string
	ImageUrl               string
}

type FindProductByIdResponse struct {
	OwnerId                string
	ProductId              string
	ProductName            string
	ProductMeasurementUnit string
	ProductPrice           float64
	IsActive               bool
	CreatedDate            string
	LastModified           string
	ImageUrl               string
}
