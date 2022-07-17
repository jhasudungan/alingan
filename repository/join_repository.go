package repository

import (
	"alingan/core/config"
	"alingan/core/model"
)

type JoinRepository interface {
	FindTransactionByOwnerId(ownerId string) ([]model.FindTransactionByOwnerIdDTO, error)
	FindAgentByOwnerId(ownerId string) ([]model.FindAgentByOwnerIdDTO, error)
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
