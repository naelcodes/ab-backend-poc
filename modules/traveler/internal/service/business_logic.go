package service

import (
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/payloads"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllTravelersService(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.Traveler], error) {
	return s.Repository.GetAll(queryParams)
}

func (s *Service) CreateTravelerService(payload payloads.CreateTravelerPayload) (*models.Traveler, error) {
	return s.Repository.Save(&payload.Traveler)
}

func (s *Service) GetTravelerByIdService(id int) (*models.Traveler, error) {
	return s.Repository.GetById(id)
}

func (s *Service) UpdateTravelerService(id int, payload payloads.UpdateTravelerPayload) (*models.Traveler, error) {
	return s.Repository.Update(id, &payload.Traveler)
}
