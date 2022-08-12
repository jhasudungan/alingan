package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportRepository(t *testing.T) {

	reportRepo := ReportRepositoryImpl{}

	t.Run("TestOwnerMostPurchasedProductByQuantity", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerMostPurchasedProductByQuantity("owner-001")

		assert.Equal(t, "PRDa543809e-f36f-443a-a815-64c0e2f0e09c", results[0].ProductId)
		assert.Equal(t, "prd-002", results[1].ProductId)
		assert.Equal(t, "prd-001", results[2].ProductId)

	})

	t.Run("TestOwnerMostPurchasedProductByRevenue", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerMostPurchasedProductByRevenue("owner-001")

		assert.Equal(t, "prd-002", results[0].ProductId)
		assert.Equal(t, "PRDa543809e-f36f-443a-a815-64c0e2f0e09c", results[1].ProductId)
		assert.Equal(t, "prd-001", results[2].ProductId)

	})

	t.Run("TestOwnerStoreWithTheMostTransaction", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerStoreWithTheMostTransaction("owner-001")

		assert.Equal(t, "str-001", results[0].StoreId)
		assert.Equal(t, "str-002", results[1].StoreId)

	})

	t.Run("TestOwnerAgentWithTheMostTransaction", func(t *testing.T) {

		results, _ := reportRepo.FindOwnerAgentWithTheMostTransaction("owner-001")

		assert.Equal(t, "agent-001", results[0].AgentId)
		assert.Equal(t, "AGT454497b9-74d1-4bb0-8753-962a962e31f6", results[1].AgentId)

	})
}
