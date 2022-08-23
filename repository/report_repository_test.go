package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportRepository(t *testing.T) {

	/**
	- run alingan-test-source-script.sql on "core" schema before run below test
	- .env need to be present in "repository" package in order to run go test
	*/

	reportRepo := ReportRepositoryImpl{}

	t.Run("TestOwnerMostPurchasedProductByQuantity", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerMostPurchasedProductByQuantity("owner-001")

		assert.Equal(t, "prd-002", results[0].ProductId)
		assert.Equal(t, "prd-001", results[1].ProductId)

	})

	t.Run("TestOwnerMostPurchasedProductByRevenue", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerMostPurchasedProductByRevenue("owner-001")

		assert.Equal(t, "prd-002", results[0].ProductId)
		assert.Equal(t, "prd-001", results[1].ProductId)

	})

	t.Run("TestOwnerStoreWithTheMostTransaction", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerStoreWithTheMostTransaction("owner-001")

		assert.Equal(t, "str-001", results[0].StoreId)

	})

	t.Run("TestOwnerAgentWithTheMostTransaction", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerAgentWithTheMostTransaction("owner-001")

		assert.Equal(t, "agent-001", results[0].AgentId)

	})
}
