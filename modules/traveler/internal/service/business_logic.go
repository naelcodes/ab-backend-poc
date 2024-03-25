package service

import "neema.co.za/rest/utils/models"

func (s *Service) GetAllTravelersService() ([]*models.Traveler, error) {
	return s.Repository.GetAll()
}
