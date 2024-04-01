package service

import (
	"fmt"

	"neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/payloads"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllCustomerService(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.Customer], error) {
	logger.Info("Getting all customers")
	return s.Repository.GetAll(queryParams)
}

func (s *Service) GetCustomerService(id int) (*models.Customer, error) {
	logger.Info("Getting customer")
	return s.Repository.GetById(id)
}

func (s *Service) CreateCustomerService(payload payloads.CreateCustomerPayload) (*models.Customer, error) {
	logger.Info(fmt.Sprintf("Creating customer: %v", payload))

	payload.Customer.IdCountry = 40
	payload.Customer.IdCurrency = 550
	// payload.Customer.IdGroup = 1
	payload.Customer.AbKey = helpers.GenerateRandomString(15)

	return s.Repository.Save(&payload.Customer)
}

func (s *Service) UpdateCustomerService(id int, payload payloads.UpdateCustomerPayload) error {
	logger.Info(fmt.Sprintf("Updating customer: %v", payload))
	return s.Repository.Update(id, &payload.Customer)
}
