package repository

import (
	"alingan/config"
	"alingan/entity"
)

type ProductRepository interface {
	Insert(data entity.Product) error
	Update(data entity.Product, productId string) error
	FindProductsByOwnerId(ownerId string) ([]entity.Product, error)
	FindProductById(productId string) (entity.Product, error)
	SetInactive(productId string) error
	SetActive(productId string) error
	CheckExist(productId string) (bool, error)
	Delete(productId string) error
}

type ProductRepositoryImpl struct{}

func (p *ProductRepositoryImpl) Insert(data entity.Product) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.product (product_id, owner_id, product_name, product_measurement_unit, product_price, is_active, created_date, last_modified) values($1, $2, $3, $4, $5, true, now(), now());"

	_, err = con.Exec(sql,
		data.ProductId,
		data.OwnerId,
		data.ProductName,
		data.ProductMeasurementUnit,
		data.ProductPrice)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) Update(data entity.Product, productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.product set product_name= $1, product_measurement_unit= $2, product_price = $3, last_modified= now() where product_id= $4"

	_, err = con.Exec(sql,
		data.ProductName,
		data.ProductMeasurementUnit,
		data.ProductPrice,
		productId)

	return nil
}

func (p *ProductRepositoryImpl) FindProductsByOwnerId(ownerId string) ([]entity.Product, error) {

	products := make([]entity.Product, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return products, err
	}

	sql := "select p.product_id, p.owner_id, p.product_name, p.product_measurement_unit, p.product_price, p.is_active, p.created_date, p.last_modified from core.product p where p.owner_id = $1 order by p.last_modified desc"

	rows, err := con.Query(sql, ownerId)

	if err != nil {
		return products, err
	}

	for rows.Next() {

		product := entity.Product{}

		err = rows.Scan(&product.ProductId,
			&product.OwnerId,
			&product.ProductName,
			&product.ProductMeasurementUnit,
			&product.ProductPrice,
			&product.IsActive,
			&product.CreatedDate,
			&product.LastModified)

		if err != nil {
			return products, err
		}

		products = append(products, product)

	}

	return products, nil
}

func (p *ProductRepositoryImpl) FindProductById(productId string) (entity.Product, error) {

	product := entity.Product{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return product, err
	}

	sql := "select p.product_id, p.owner_id, p.product_name, p.product_measurement_unit, p.product_price, p.is_active, p.created_date, p.last_modified from core.product p where p.product_id = $1"

	row := con.QueryRow(sql, productId)

	err = row.Scan(&product.ProductId,
		&product.OwnerId,
		&product.ProductName,
		&product.ProductMeasurementUnit,
		&product.ProductPrice,
		&product.IsActive,
		&product.CreatedDate,
		&product.LastModified)

	if err != nil {
		return product, err
	}

	return product, nil

}

func (p *ProductRepositoryImpl) SetInactive(productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.product set is_active = false where product_id = $1"

	_, err = con.Exec(sql, productId)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) SetActive(productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.product set is_active = true where product_id = $1"

	_, err = con.Exec(sql, productId)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) CheckExist(productId string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.product p where p.product_id = $1)"

	row := con.QueryRow(sql, productId)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *ProductRepositoryImpl) Delete(productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.product where product_id= $1"

	_, err = con.Exec(sql, productId)

	if err != nil {
		return err
	}

	return nil
}
