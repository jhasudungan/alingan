package repository

import (
	"alingan/core/config"
	"time"
)

type JoinRepository interface {
	FindTransactionByOwnerId(ownerId string) ([]map[string]interface{}, error)
}

type JoinRepositoryImpl struct{}

func (j *JoinRepositoryImpl) FindTransactionByOwnerId(ownerId string) ([]map[string]interface{}, error) {

	results := make([]map[string]interface{}, 0)

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

		transactionId := ""
		transactionDate := time.Now()
		storeName := ""
		storeId := ""
		agentId := ""
		agentName := ""
		transactionTotal := float64(0)

		err = rows.Scan(
			&transactionId,
			&transactionDate,
			&storeName,
			&storeId,
			&agentId,
			&agentName,
			&transactionTotal)

		data := make(map[string]interface{})
		data["transactionId"] = transactionId
		data["transactionDate"] = transactionDate
		data["storeName"] = storeName
		data["storeId"] = storeId
		data["agentId"] = agentId
		data["agentName"] = agentName
		data["transactionTotal"] = transactionTotal

		results = append(results, data)
	}

	return results, nil
}
