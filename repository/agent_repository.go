package repository

import (
	"alingan/config"
	"alingan/entity"
)

type AgentRepository interface {
	Insert(data entity.Agent) error
	Update(data entity.Agent, agentId string) error
	FindAgentById(agentId string) (entity.Agent, error)
	FindAgentByEmail(agentEmail string) (entity.Agent, error)
	SetInactive(agentId string) error
	SetActive(agentId string) error
	CheckEmailExist(agentEmail string) (bool, error)
	CheckExist(agentId string) (bool, error)
	Delete(agentId string) error
}

type AgentRepositoryImpl struct{}

func (a *AgentRepositoryImpl) Insert(data entity.Agent) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "insert into core.agent (agent_id, store_id, agent_name, agent_email, agent_password, is_active, created_date, last_modified) values($1, $2, $3, $4, $5, true, now(), now())"

	_, err = con.Exec(sql,
		data.AgentId,
		data.StoreId,
		data.AgentName,
		data.AgentEmail,
		data.AgentPassword)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentRepositoryImpl) Update(data entity.Agent, agentId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.agent set agent_name=$1, agent_email=$2, agent_password=$3, last_modified = now() where agent_id=$4"

	_, err = con.Exec(sql,
		data.AgentName,
		data.AgentEmail,
		data.AgentPassword,
		agentId)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentRepositoryImpl) FindAgentById(agentId string) (entity.Agent, error) {

	agent := entity.Agent{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return agent, err
	}

	sql := "select a.agent_id, a.store_id, a.agent_name, a.agent_email, a.agent_password, a.is_active, a.created_date, a.last_modified from core.agent a where a.agent_id = $1"

	row := con.QueryRow(sql, agentId)

	err = row.Scan(
		&agent.AgentId,
		&agent.StoreId,
		&agent.AgentName,
		&agent.AgentEmail,
		&agent.AgentPassword,
		&agent.IsActive,
		&agent.CreatedDate,
		&agent.LastModified)

	if err != nil {
		return agent, err
	}

	return agent, nil

}

func (a *AgentRepositoryImpl) FindAgentByEmail(agentEmail string) (entity.Agent, error) {

	agent := entity.Agent{}

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return agent, err
	}

	sql := "select a.agent_id, a.store_id, a.agent_name, a.agent_email, a.agent_password, a.is_active, a.created_date, a.last_modified from core.agent a where a.agent_email = $1"

	row := con.QueryRow(sql, agentEmail)

	err = row.Scan(
		&agent.AgentId,
		&agent.StoreId,
		&agent.AgentName,
		&agent.AgentEmail,
		&agent.AgentPassword,
		&agent.IsActive,
		&agent.CreatedDate,
		&agent.LastModified)

	if err != nil {
		return agent, err
	}

	return agent, nil
}

func (a *AgentRepositoryImpl) SetInactive(agentId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.agent set is_active = false where agent_id = $1"

	_, err = con.Exec(sql, agentId)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentRepositoryImpl) SetActive(agentId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "update core.agent set is_active = true where agent_id = $1"

	_, err = con.Exec(sql, agentId)

	if err != nil {
		return err
	}

	return nil
}

func (a *AgentRepositoryImpl) CheckExist(agentId string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.agent p where p.agent_id = $1)"

	row := con.QueryRow(sql, agentId)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (a *AgentRepositoryImpl) CheckEmailExist(agentEmail string) (bool, error) {

	result := false

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := "select exists (select 1 from core.agent p where p.agent_email = $1)"

	row := con.QueryRow(sql, agentEmail)

	err = row.Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (a *AgentRepositoryImpl) Delete(agentId string) error {

	con, err := config.CreateDBConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := "delete from core.agent where agent_id= $1"

	_, err = con.Exec(sql, agentId)

	if err != nil {
		return err
	}

	return nil
}
