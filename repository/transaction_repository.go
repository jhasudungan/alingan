package repository

import (
	"alingan/core/config"
	"alingan/core/entity"
)

type TransactionRepository interface {
	Insert(data entity.Transaction) error
	FindById(transactionId string) (entity.Transaction, error)
}

type TransactionRepositoryImpl struct{}

func (t *TransactionRepositoryImpl) Insert(data entity.Transaction) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.transaction (transaction_id, transaction_date, agent_id, store_id, transaction_total, created_date, last_modified) values($1, $2, $3, $4, $5, now(), now());"

	_, err = con.Exec(sql,
		data.TransactionId,
		data.TransactionDate,
		data.AgentId,
		data.StoreId,
		data.TransactionTotal)

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryImpl) FindById(transactionId string) (entity.Transaction, error) {

	transaction := entity.Transaction{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return transaction, err
	}

	sql := "select t.transaction_id, t.transaction_date, t.agent_id, t.store_id, t.transaction_total, t.created_date, t.last_modified from core.transaction t where t.transaction_id = $1"

	row := con.QueryRow(sql, transactionId)

	err = row.Scan(
		&transaction.TransactionId,
		&transaction.TransactionDate,
		&transaction.AgentId,
		&transaction.StoreId,
		&transaction.TransactionTotal,
		&transaction.CreatedDate,
		&transaction.LastModified)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
