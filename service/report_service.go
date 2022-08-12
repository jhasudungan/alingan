package service

import (
	"alingan/model"
	"alingan/repository"
)

type ReportService interface {
	WebDashboardReport(ownerId string) (model.WebDashboardReportResponse, error)
}

type ReportServiceImpl struct {
	ReportRepository repository.ReportRepository
}

func (r *ReportServiceImpl) WebDashboardReport(ownerId string) (model.WebDashboardReportResponse, error) {

	result := model.WebDashboardReportResponse{}

	data1, err := r.ReportRepository.FindOwnerMostPurchasedProductByQuantity(ownerId)

	if err != nil {
		return result, err
	}

	data2, err := r.ReportRepository.FindOwnerMostPurchasedProductByRevenue(ownerId)

	if err != nil {
		return result, err
	}

	data3, err := r.ReportRepository.FindOwnerAgentWithTheMostTransaction(ownerId)

	if err != nil {
		return result, err
	}

	data4, err := r.ReportRepository.FindOwnerStoreWithTheMostTransaction(ownerId)

	if err != nil {
		return result, err
	}

	result.OwnerMostPurchasedProductByQuantity = data1
	result.OwnerMostPurchasedProductByRevenue = data2
	result.OwnerAgentWithTheMostTransaction = data3
	result.OwnerStoreWithTheMostTransaction = data4

	return result, nil

}
