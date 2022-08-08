package repository

import (
	"alingan/config"
	"alingan/model"
)

type JoinRepository interface {
	FindTransactionByOwnerId(ownerId string) ([]model.FindTransactionByOwnerIdDTO, error)
	FindTransactionAgentAndStoreByTransactionId(transactionId string) (model.FindTransactionAgentAndStoreByTransactionIdDTO, error)
	FindTransactionItemAndProductByTransactionId(transactionId string) ([]model.FindTransactionItemAndProductByTransactionIdDTO, error)
	FindAgentByOwnerId(ownerId string) ([]model.FindAgentByOwnerIdDTO, error)
	FindOwnerByAgentId(agentId string) (model.FindOwnerByAgentIdDTO, error)
}

type JoinRepositoryImpl struct{}

func (j *JoinRepositoryImpl) FindTransactionByOwnerId(ownerId string) ([]model.FindTransactionByOwnerIdDTO, error) {

	results := make([]model.FindTransactionByOwnerIdDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select t.transaction_id , t.transaction_date , s.store_name , s.store_id, t.agent_id, a.agent_name , t.transaction_total  from core.transaction t inner join core.agent a on t.agent_id = a.agent_id inner join core.store s on t.store_id = s.store_id where s.owner_id = $1 order by t.transaction_date desc"

	rows, err := con.Query(
		sql,
		ownerId)

	if err != nil {
		return results, err
	}

	for rows.Next() {

		transaction := model.FindTransactionByOwnerIdDTO{}

		err = rows.Scan(
			&transaction.TransactionId,
			&transaction.TransactionDate,
			&transaction.StoreName,
			&transaction.StoreId,
			&transaction.AgentId,
			&transaction.AgentName,
			&transaction.TransactionTotal)

		results = append(results, transaction)
	}

	return results, nil
}

func (j *JoinRepositoryImpl) FindTransactionAgentAndStoreByTransactionId(transactionId string) (model.FindTransactionAgentAndStoreByTransactionIdDTO, error) {

	result := model.FindTransactionAgentAndStoreByTransactionIdDTO{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select t.transaction_id , t.transaction_date , s.store_name , s.store_id, t.agent_id, a.agent_name , t.transaction_total from core.transaction t inner join core.agent a on t.agent_id = a.agent_id inner join core.store s on t.store_id = s.store_id where t.transaction_id = $1 order by t.transaction_date desc"

	row := con.QueryRow(
		sql,
		transactionId)

	err = row.Scan(
		&result.TransactionId,
		&result.TransactionDate,
		&result.StoreName,
		&result.StoreId,
		&result.AgentId,
		&result.AgentName,
		&result.TransactionTotal)

	return result, nil
}

func (j *JoinRepositoryImpl) FindTransactionItemAndProductByTransactionId(transactionId string) ([]model.FindTransactionItemAndProductByTransactionIdDTO, error) {

	results := make([]model.FindTransactionItemAndProductByTransactionIdDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select ti.transaction_item_id , ti.transaction_id , p.product_id, p.product_name , ti.used_price , ti.buy_quantity from core.transaction_item ti inner join core.product p on ti.product_id = p.product_id where ti.transaction_id  = $1"

	rows, err := con.Query(sql, transactionId)

	for rows.Next() {

		data := model.FindTransactionItemAndProductByTransactionIdDTO{}

		err = rows.Scan(
			&data.TransactionItemId,
			&data.TransactionId,
			&data.ProductId,
			&data.ProductName,
			&data.UsedPrice,
			&data.BuyQuantity)

		if err != nil {
			return results, err
		}

		results = append(results, data)
	}

	return results, nil
}

func (j *JoinRepositoryImpl) FindAgentByOwnerId(ownerId string) ([]model.FindAgentByOwnerIdDTO, error) {

	results := make([]model.FindAgentByOwnerIdDTO, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return results, err
	}

	sql := "select a.agent_id , a.agent_name , a.agent_email , s.store_id , s.store_name , a.is_active from core.agent a inner join core.store s on a.store_id = s.store_id inner join core.owner o on s.owner_id = o.owner_id where o.owner_id = $1"

	rows, err := con.Query(
		sql,
		ownerId)

	if err != nil {
		return results, err
	}

	for rows.Next() {

		agent := model.FindAgentByOwnerIdDTO{}

		err = rows.Scan(
			&agent.AgentId,
			&agent.AgentName,
			&agent.AgentEmail,
			&agent.StoreId,
			&agent.StoreName,
			&agent.IsActive)

		results = append(results, agent)
	}

	return results, nil
}

func (j *JoinRepositoryImpl) FindOwnerByAgentId(agentId string) (model.FindOwnerByAgentIdDTO, error) {

	result := model.FindOwnerByAgentIdDTO{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select o.owner_id , o.owner_name , o.owner_email , o.is_active, o.created_date , o.last_modified from core.owner o inner join core.store s on o.owner_id = s.owner_id inner join core.agent a on s.store_id = a.store_id where a.agent_id = $1"

	row := con.QueryRow(sql, agentId)

	err = row.Scan(&result.OwnerId,
		&result.OwnerName,
		&result.OwnerEmail,
		&result.IsActive,
		&result.CreatedDate,
		&result.LastModified)

	if err != nil {
		return result, err
	}

	return result, nil

}
