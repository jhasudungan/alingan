package repository

import "alingan/config"

// Below query only used for cleaning up data after testing , is not executed in production

type TestingRepository struct{}

func (t *TestingRepository) DeleteAllStoreByOwner(ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.store where owner_id = $1"

	_, err = con.Exec(sql, ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteAllProductByOwner(ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.product where owner_id = $1"

	_, err = con.Exec(sql, ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteAllAgentByStore(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.agent where store_id = $1"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteAllTransactionByStore(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.transaction where store_id = $1"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteAllTransactionItemByStore(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.transaction_item where transaction_id in (select transaction_id from core.transaction where store_id = $1)"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteAllProductImageByProductId(productId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.product_image where product_id = $1"

	_, err = con.Exec(sql, productId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteTransactionById(transactionId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.transaction where transaction_id = $1"

	_, err = con.Exec(sql, transactionId)

	if err != nil {
		return err
	}

	return nil
}

func (t *TestingRepository) DeleteTransactionItemByTransactionId(transactionId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.transaction_item where transaction_id = $1"

	_, err = con.Exec(sql, transactionId)

	if err != nil {
		return err
	}

	return nil
}
