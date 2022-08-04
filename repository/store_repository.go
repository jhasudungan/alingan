package repository

import (
	"alingan/config"
	"alingan/entity"
)

type StoreRepository interface {
	Insert(data entity.Store) error
	Update(data entity.Store, storeId string) error
	FindStoresByOwnerId(ownerId string) ([]entity.Store, error)
	FindStoreById(storeId string) (entity.Store, error)
	SetInactive(storeId string) error
	SetActive(storeId string) error
	CheckExist(storeId string) (bool, error)
	Delete(storeId string) error
}

type StoreRepositoryImpl struct{}

func (s *StoreRepositoryImpl) Insert(data entity.Store) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.store (store_id, owner_id, store_name, store_address, is_active, created_date, last_modified) values($1, $2, $3, $4, true, now(), now())"

	_, err = con.Exec(sql, data.StoreId, data.OwnerId, data.StoreName, data.StoreAddress)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreRepositoryImpl) Update(data entity.Store, storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.store set store_name = $1, store_address = $2, last_modified = now() where store_id = $3"

	_, err = con.Exec(sql, data.StoreName, data.StoreAddress, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreRepositoryImpl) FindStoresByOwnerId(ownerId string) ([]entity.Store, error) {

	stores := make([]entity.Store, 0)

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return stores, err
	}

	sql := "select s.store_id, s.owner_id, s.store_name, s.store_address, s.is_active, s.created_date, s.last_modified from core.store s where s.owner_id = $1 order by created_date desc"

	rows, err := con.Query(sql, ownerId)

	if err != nil {
		return stores, err
	}

	for rows.Next() {

		store := entity.Store{}

		err := rows.Scan(&store.StoreId,
			&store.OwnerId,
			&store.StoreName,
			&store.StoreAddress,
			&store.IsActive,
			&store.CreatedDate,
			&store.LastModified)

		if err != nil {
			return stores, err
		}

		stores = append(stores, store)

	}

	return stores, nil
}

func (s *StoreRepositoryImpl) FindStoreById(storeId string) (entity.Store, error) {

	store := entity.Store{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return store, err
	}

	sql := "select s.store_id, s.owner_id, s.store_name, s.store_address, s.is_active, s.created_date, s.last_modified from core.store s where s.store_id = $1"

	row := con.QueryRow(sql, storeId)

	err = row.Scan(
		&store.StoreId,
		&store.OwnerId,
		&store.StoreName,
		&store.StoreAddress,
		&store.IsActive,
		&store.CreatedDate,
		&store.LastModified)

	if err != nil {
		return store, err
	}

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

func (s *StoreRepositoryImpl) SetActive(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.store set is_active = true where store_id = $1"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (s *StoreRepositoryImpl) CheckExist(storeId string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.store p where p.store_id = $1)"

	row := con.QueryRow(sql, storeId)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

func (s *StoreRepositoryImpl) Delete(storeId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.store where store_id= $1"

	_, err = con.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}
