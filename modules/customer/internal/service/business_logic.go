package service

import (
	"fmt"

	"neema.co.za/rest/modules/customer/internal/domain"

	"neema.co.za/rest/utils/dto"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllCustomerService(queryParams *types.GetQueryParams) (*dto.GetAllCustomersDTO, error) {
	logger.Info("Getting all customers")
	return s.Repository.GetAll(queryParams)
}

func (s *Service) GetCustomerService(id int) (*dto.GetCustomerDTO, error) {
	logger.Info("Getting customer")
	return s.Repository.GetById(id)
}

func (s *Service) CreateCustomerService(createCustomerDTO *dto.CreateCustomerDTO) (*dto.GetCustomerDTO, error) {
	logger.Info(fmt.Sprintf("Creating customer: %v", createCustomerDTO))
	customer := domain.NewCustomerBuilder().
		SetCustomerName(createCustomerDTO.CustomerName).
		SetAlias(createCustomerDTO.Alias).
		SetAbKey().
		SetTmcClientNumber(createCustomerDTO.TmcClientNumber).
		SetAccountNumber(createCustomerDTO.AccountNumber).
		SetState(createCustomerDTO.State).
		Build()

	return s.Repository.Save(customer)
}
