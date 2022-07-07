package util

import (
	"log"
	"testing"
)

func TestServiceUtil(t *testing.T) {

	t.Run("TestGenerateId", func(t *testing.T) {

		serviceUtil := &ServiceUtil{}
		id := serviceUtil.GenerateId("STR")
		log.Printf("Id : %v", id)

	})
}
