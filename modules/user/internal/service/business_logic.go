package service

import (
	"neema.co.za/rest/utils/models"
)

func (s *Service) GetUserByID(id int) (*models.User, error) {
	// Business logic (if any)
	return s.Repository.GetUserByID(id)
}
