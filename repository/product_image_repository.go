package repository

import (
	"alingan/config"
	"alingan/entity"
)

type ProductImageRepository interface {
	Insert(data entity.ProductImage) error
	FindProductImageById(productImageId string) (entity.ProductImage, error)
	FindProductImageByProductId(productId string) ([]entity.ProductImage, error)
	Delete(productId string) error
}

type ProductImageRepositoryImpl struct{}

func (p *ProductImageRepositoryImpl) Insert(data entity.ProductImage) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.product_image (product_image_id, product_id, location_path, created_date, last_modified) values($1, $2, $3, now(), now())"

	_, err = con.Exec(sql,
		data.ProductImageId,
		data.ProductId,
		data.LocationPath)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductImageRepositoryImpl) Delete(productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.product_image where product_image_id= $1"

	_, err = con.Exec(sql, productId)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductImageRepositoryImpl) FindProductImageById(productImageId string) (entity.ProductImage, error) {

	productImage := entity.ProductImage{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return productImage, err
	}

	sql := "select product_image_id, product_id, location_path, created_date, last_modified from core.product_image where product_image_id = $1"

	row := con.QueryRow(sql, productImageId)

	err = row.Scan(&productImage.ProductImageId,
		&productImage.ProductId,
		&productImage.LocationPath,
		&productImage.CreatedDate,
		&productImage.LastModified)

	if err != nil {
		return productImage, err
	}

	return productImage, nil
}

func (p *ProductImageRepositoryImpl) FindProductImageByProductId(productId string) ([]entity.ProductImage, error) {

	productImages := make([]entity.ProductImage, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return productImages, err
	}

	sql := "select product_image_id, product_id, location_path, created_date, last_modified from core.product_image where product_id = $1 order by last_modified desc;"

	rows, err := con.Query(sql, productId)

	if err != nil {
		return productImages, err
	}

	for rows.Next() {

		productImage := entity.ProductImage{}

		err = rows.Scan(&productImage.ProductImageId,
			&productImage.ProductId,
			&productImage.LocationPath,
			&productImage.CreatedDate,
			&productImage.LastModified)

		if err != nil {
			return productImages, err
		}

		productImages = append(productImages, productImage)
	}

	return productImages, nil
}
