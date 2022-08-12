package repository

import (
	"alingan/config"
	"alingan/model"
)

type ReportRepository interface {
	FindOwnerMostPurchasedProductByQuantity(ownerId string) ([]model.FindOwnerMostPurchasedProductByQuantityDTO, error)
	FindOwnerMostPurchasedProductByRevenue(ownerId string) ([]model.FindOwnerMostPurchasedProductByRevenueDTO, error)
	FindOwnerStoreWithTheMostTransaction(ownerId string) ([]model.FindOwnerStoreWithTheMostTransactionDTO, error)
	FindOwnerAgentWithTheMostTransaction(ownerId string) ([]model.FindOwnerAgentWithTheMostTransactionDTO, error)
}

type ReportRepositoryImpl struct{}

func (r *ReportRepositoryImpl) FindOwnerMostPurchasedProductByQuantity(ownerId string) ([]model.FindOwnerMostPurchasedProductByQuantityDTO, error) {

	results := make([]model.FindOwnerMostPurchasedProductByQuantityDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select sum(ti.buy_quantity) as total_purchased, p.product_id, p.product_name from core.owner o inner join core.store s on o.owner_id = s.owner_id inner join core.transaction t on s.store_id = t.store_id inner join core.transaction_item ti on t.transaction_id = ti.transaction_id inner join core.product p on ti.product_id = p.product_id where o.owner_id = $1 group  by p.product_id , p.product_name order by total_purchased desc limit 3"

	rows, err := con.Query(sql, ownerId)

	for rows.Next() {

		data := model.FindOwnerMostPurchasedProductByQuantityDTO{}

		err = rows.Scan(
			&data.TotalPurchased,
			&data.ProductId,
			&data.ProductName)

		results = append(results, data)

	}

	return results, nil
}

func (r *ReportRepositoryImpl) FindOwnerMostPurchasedProductByRevenue(ownerId string) ([]model.FindOwnerMostPurchasedProductByRevenueDTO, error) {

	results := make([]model.FindOwnerMostPurchasedProductByRevenueDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select sum(ti.buy_quantity*ti.used_price) as total_revenue, p.product_id, p.product_name from core.owner o inner join core.store s on o.owner_id = s.owner_id inner join core.transaction t on s.store_id = t.store_id inner join core.transaction_item ti on t.transaction_id = ti.transaction_id inner join core.product p on ti.product_id = p.product_id where  o.owner_id = $1 group  by p.product_id , p.product_name order by total_revenue desc limit 3"

	rows, err := con.Query(sql, ownerId)

	for rows.Next() {

		data := model.FindOwnerMostPurchasedProductByRevenueDTO{}

		err = rows.Scan(
			&data.TotalRevenue,
			&data.ProductId,
			&data.ProductName)

		results = append(results, data)

	}

	return results, nil
}

func (r *ReportRepositoryImpl) FindOwnerStoreWithTheMostTransaction(ownerId string) ([]model.FindOwnerStoreWithTheMostTransactionDTO, error) {

	results := make([]model.FindOwnerStoreWithTheMostTransactionDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select s.store_id , s.store_name , count(t.transaction_id) as store_total_transaction from core.owner o inner join core.store s on o.owner_id = s.owner_id inner join core.transaction t on t.store_id = s.store_id where o.owner_id = $1 group by s.store_id , s.store_name order by store_total_transaction desc limit 3;"

	rows, err := con.Query(sql, ownerId)

	for rows.Next() {

		data := model.FindOwnerStoreWithTheMostTransactionDTO{}

		err = rows.Scan(
			&data.StoreId,
			&data.StoreName,
			&data.StoreTotalTransaction)

		results = append(results, data)

	}

	return results, nil
}

func (r *ReportRepositoryImpl) FindOwnerAgentWithTheMostTransaction(ownerId string) ([]model.FindOwnerAgentWithTheMostTransactionDTO, error) {

	results := make([]model.FindOwnerAgentWithTheMostTransactionDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select count(t.transaction_id) as total_transaction_handled, a.agent_id, a.agent_name from core.owner o inner join core.store s on o.owner_id = s.owner_id inner join core.transaction t on s.store_id = t.store_id inner join core.agent a on t.agent_id  = a.agent_id where  o.owner_id = $1 group  by a.agent_id , a.agent_name order by total_transaction_handled desc limit 3"

	rows, err := con.Query(sql, ownerId)

	for rows.Next() {

		data := model.FindOwnerAgentWithTheMostTransactionDTO{}

		err = rows.Scan(
			&data.TotalTransactionHandled,
			&data.AgentId,
			&data.AgentName)

		results = append(results, data)

	}

	return results, nil
}
