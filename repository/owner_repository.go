package repository

import (
	"alingan/config"
	"alingan/entity"
)

type OwnerRepository interface {
	Insert(data entity.Owner) error
	Update(data entity.Owner, ownerId string) error
	UpdatePassword(data entity.Owner, ownerId string) error
	FindById(ownerId string) (entity.Owner, error)
	FindByOwnerEmail(ownerId string) (entity.Owner, error)
	CheckExist(ownerId string) (bool, error)
	CheckEmailExist(ownerEmail string) (bool, error)
	SetInactive(ownerId string) error
	Delete(ownerId string) error
}

type OwnerRepositoryImpl struct{}

func (o *OwnerRepositoryImpl) Insert(data entity.Owner) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.owner (owner_id, owner_name, owner_type, owner_email, password, is_active, created_date, last_modified) values($1,$2,$3,$4,$5,true, now(), now())"

	_, err = con.Exec(sql,
		data.OwnerId,
		data.OwnerName,
		data.OwnerType,
		data.OwnerEmail,
		data.Password)

	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepositoryImpl) Update(data entity.Owner, ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.owner set owner_name=$1, owner_type=$2, last_modified= now() WHERE owner_id=$3"

	_, err = con.Exec(sql,
		data.OwnerName,
		data.OwnerType,
		ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepositoryImpl) UpdatePassword(data entity.Owner, ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.owner set password = $1 last_modified= now() WHERE owner_id=$2"

	_, err = con.Exec(sql,
		data.Password,
		ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepositoryImpl) FindById(ownerId string) (entity.Owner, error) {

	owner := entity.Owner{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return owner, err
	}

	sql := "select owner_id, owner_name, owner_type, owner_email, password, is_active, created_date, last_modified from core.owner where owner_id = $1"

	row := con.QueryRow(sql, ownerId)

	err = row.Scan(
		&owner.OwnerId,
		&owner.OwnerName,
		&owner.OwnerType,
		&owner.OwnerEmail,
		&owner.Password,
		&owner.IsActive,
		&owner.CreatedDate,
		&owner.LastModified)

	if err != nil {
		return owner, err
	}

	return owner, nil
}

func (o *OwnerRepositoryImpl) FindByOwnerEmail(ownerId string) (entity.Owner, error) {

	owner := entity.Owner{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return owner, err
	}

	sql := "select owner_id, owner_name, owner_type, owner_email, password, is_active, created_date, last_modified from core.owner where owner_email = $1"

	row := con.QueryRow(sql, ownerId)

	err = row.Scan(
		&owner.OwnerId,
		&owner.OwnerName,
		&owner.OwnerType,
		&owner.OwnerEmail,
		&owner.Password,
		&owner.IsActive,
		&owner.CreatedDate,
		&owner.LastModified)

	if err != nil {
		return owner, err
	}

	return owner, nil
}

func (o *OwnerRepositoryImpl) CheckExist(ownerId string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.owner o where o.owner_id = $1)"

	row := con.QueryRow(sql, ownerId)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

func (o *OwnerRepositoryImpl) CheckEmailExist(ownerEmail string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.owner o where o.owner_email = $1)"

	row := con.QueryRow(sql, ownerEmail)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

func (o *OwnerRepositoryImpl) SetInactive(ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.owner set is_active = false where owner_id = $1"

	_, err = con.Exec(sql, ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepositoryImpl) Delete(ownerId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.owner where owner_id = $1"

	_, err = con.Exec(sql, ownerId)

	return nil

}
