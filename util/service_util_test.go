package util

import (
	"log"
	"testing"
)

func TestServiceUtil(t *testing.T) {

	t.Run("TestGenerateId", func(t *testing.T) {

		serviceUtil := &ServiceUtil{}

		id, err := serviceUtil.GenerateId("STR")

		if err != nil {
			log.Fatal("Error Test : " + err.Error())
			t.FailNow()
		}

		log.Printf("Id : %v", id)

	})
}
