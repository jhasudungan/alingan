package repository

import "alingan/core/config"

// Below query only used for cleaning up data after testing

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

func (t *TestingRepository) DeleteAllProductByOwnerA(ownerId string) error {

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

	sql := "delete from core.product where store_id = $1"

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

func (t *TestingRepository) DeleteAllTransactionItemByTransaction(transactionId string) error {

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
