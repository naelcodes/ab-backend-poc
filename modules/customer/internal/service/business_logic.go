package service

import (
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"
)

func (s *Service) GetAllCustomerService() ([]*models.Customer, error) {
	logger.Info("Getting all customers")
	return s.Repository.GetAll()
}
