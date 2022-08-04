package repository

import (
	"alingan/config"
	"alingan/entity"
)

type TransactionReceiptRepository interface {
	Insert(data entity.TransactionReceipt) error
	FindTransactionReceiptByTransactionId(transactionId string) ([]entity.TransactionReceipt, error)
	FindTransactionReceiptById(transactionReceiptId string) (entity.TransactionReceipt, error)
	Delete(transactionReceiptId string) error
}

type TransactionReceiptRepositoryImpl struct{}

func (t *TransactionReceiptRepositoryImpl) Insert(data entity.TransactionReceipt) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.transaction_receipt (transaction_receipt_id, transaction_id, location_path, created_date, last_modified) values($1, $2, $3, now(), now())"

	_, err = con.Exec(sql, data.TransactionReceiptId, data.TransactionId, data.LocationPath)

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionReceiptRepositoryImpl) FindTransactionReceiptByTransactionId(transactionId string) ([]entity.TransactionReceipt, error) {

	transactionReceipts := make([]entity.TransactionReceipt, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return transactionReceipts, err
	}

	sql := "select transaction_receipt_id, transaction_id, location_path, created_date, last_modified from core.transaction_receipt where transaction_id = $1 order by created_date desc"

	rows, err := con.Query(sql, transactionId)

	if err != nil {
		return transactionReceipts, err
	}

	for rows.Next() {

		transactionReceipt := entity.TransactionReceipt{}

		err = rows.Scan(
			&transactionReceipt.TransactionReceiptId,
			&transactionReceipt.TransactionId,
			&transactionReceipt.LocationPath,
			&transactionReceipt.CreatedDate,
			&transactionReceipt.LastModified)

		if err != nil {
			return transactionReceipts, err
		}

		transactionReceipts = append(transactionReceipts, transactionReceipt)

	}

	return transactionReceipts, nil

}

func (t *TransactionReceiptRepositoryImpl) FindTransactionReceiptById(transactionReceiptId string) (entity.TransactionReceipt, error) {

	transactionReceipt := entity.TransactionReceipt{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return transactionReceipt, err
	}

	sql := "select transaction_receipt_id, transaction_id, location_path, created_date, last_modified from core.transaction_receipt where transaction_receipt_id = $1 order by created_date desc"

	row := con.QueryRow(sql, transactionReceiptId)

	err = row.Scan(
		&transactionReceipt.TransactionReceiptId,
		&transactionReceipt.TransactionId,
		&transactionReceipt.LocationPath,
		&transactionReceipt.CreatedDate,
		&transactionReceipt.LastModified)

	if err != nil {
		return transactionReceipt, err
	}

	return transactionReceipt, nil
}

func (t *TransactionReceiptRepositoryImpl) Delete(transactionReceiptId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.transaction_receipt where transaction_receipt_id = $1"

	_, err = con.Exec(sql, transactionReceiptId)

	if err != nil {
		return err
	}

	return nil

}
