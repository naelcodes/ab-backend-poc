package service

import (
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllTravelersService(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.Traveler], error) {
	return s.Repository.GetAll(queryParams)
}
