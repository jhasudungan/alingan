package repository

import (
	"alingan/core/config"
	"alingan/core/entity"
)

type TransactionItemRepository interface {
	Insert(data entity.TransactionItem) error
	FindByTransactionId(transactionId string) ([]entity.TransactionItem, error)
}

type TransactionItemRepositoryImpl struct{}

func (t *TransactionItemRepositoryImpl) Insert(data entity.TransactionItem) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.transaction_item (transaction_item_id, product_id, transaction_id, used_price, buy_quantity, created_date, last_modified) values($1, $2, $3, $4, $5, now(), now());"

	_, err = con.Exec(sql,
		data.TransactionItemId,
		data.ProductId,
		data.TransactionId,
		data.UsedPrice,
		data.BuyQuantity)

	if err != nil {
		return err
	}

	return nil

}

func (t *TransactionItemRepositoryImpl) FindByTransactionId(transactionId string) ([]entity.TransactionItem, error) {

	items := make([]entity.TransactionItem, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return items, err
	}

	sql := "select t.transaction_item_id, t.product_id, t.transaction_id, t.used_price, t.buy_quantity, t.created_date, t.last_modified from core.transaction_item t where t.transaction_id = $1"

	rows, err := con.Query(sql, transactionId)

	if err != nil {
		return items, err
	}

	for rows.Next() {

		data := entity.TransactionItem{}

		err = rows.Scan(
			&data.TransactionItemId,
			&data.ProductId,
			&data.TransactionId,
			&data.UsedPrice,
			&data.BuyQuantity,
			&data.CreatedDate,
			&data.LastModified)

		items = append(items, data)

	}

	return items, nil
}
