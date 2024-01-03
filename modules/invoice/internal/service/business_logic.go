package service

import (
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllInvoiceService(queryParams *types.GetQueryParams) (*types.GetAllDTO[any], error) {
	logger.Info("Getting all invoices")
	return s.Repository.GetAll(queryParams)

}

func (s *Service) GetInvoiceService(id int, queryParams *types.GetQueryParams) (any, error) {
	logger.Info("Getting invoice")
	return s.Repository.GetById(id, queryParams)
}
