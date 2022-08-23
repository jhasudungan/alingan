package config

import (
	"log"
	"testing"
)

func TestCreateDbConnection(t *testing.T) {

	con, err := CreateDBConnection()

	if err != nil {
		log.Fatalf("Error Test Connection %v ", err.Error())
		t.FailNow()
	}

	err = con.Ping()

	if err != nil {
		log.Fatalf("Error Test Connection %v ", err.Error())
		t.FailNow()
	}

	log.Print("Success create connection ")

}
