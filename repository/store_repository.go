package repository

import (
	"alingan/core/config"
	"time"
)

type StoreRepository interface {
	Insert(data map[string]interface{}) error
	Update(data map[string]interface{}) error
	FindStoresByOwnerId(ownerId string) ([]map[string]interface{}, error)
	FindStoreById(storeId string) (map[string]interface{}, error)
	SetInactive(storeId string) error
}

type StoreRepositoryImpl struct{}

func (s *StoreRepositoryImpl) Insert(data map[string]interface{}) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.store (store_id, owner_id, store_name, store_address, is_active, created_date, last_modified) values($1, $2, $3, $4, true, now(), now())"

	_, err = con.Exec(sql, data["storeId"], data["ownerId"], data["storeName"], data["storeAddress"])

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreRepositoryImpl) Update(data map[string]interface{}, storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.store set store_name = $1, store_address = $2, last_modified = now() where store_id = $3"

	_, err = con.Exec(sql, data["storeName"], data["storeAddress"], storeId)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreRepositoryImpl) FindStoresByOwnerId(ownerId string) ([]map[string]interface{}, error) {

	var stores []map[string]interface{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return stores, err
	}

	sql := "select s.store_id, s.store_name , s.is_active from core.store s where s.owner_id = $1 order by s.store_name desc"

	rows, err := con.Query(sql, ownerId)

	if err != nil {
		return stores, err
	}

	for rows.Next() {

		var storeId string
		var storeName string
		var isActive bool

		err := rows.Scan(&storeId, &storeName, &isActive)

		if err != nil {
			return stores, err
		}

		store := make(map[string]interface{})
		store["storeId"] = storeId
		store["storeName"] = storeName
		store["isActive"] = isActive

		stores = append(stores, store)

	}

	return stores, nil
}

func (s *StoreRepositoryImpl) FindStoreById(storeId string) (map[string]interface{}, error) {

	var store = make(map[string]interface{})

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return store, err
	}

	sql := "select s.store_id, s.owner_id, s.store_name, s.store_address, s.is_active, s.created_date, s.last_modified from core.store s where s.store_id = $1"

	row := con.QueryRow(sql, storeId)

	var ownerId string
	var storeName string
	var storeAddress string
	var isActive bool
	var createdDate time.Time
	var lastModified time.Time

	err = row.Scan(
		&storeId,
		&ownerId,
		&storeName,
		&storeAddress,
		&isActive,
		&createdDate,
		&lastModified)

	if err != nil {
		return store, err
	}

	store["storeId"] = storeId
	store["ownerId"] = ownerId
	store["storeName"] = storeName
	store["storeAddress"] = storeAddress
	store["isActive"] = isActive
	store["createdDate"] = createdDate
	store["lastModified"] = lastModified

	return store, nil

}

func (s *StoreRepositoryImpl) SetInactive(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.store set is_active = false where store_id = $1"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}
