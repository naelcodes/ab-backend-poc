package service

import (
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllTravelItemsService__(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.TravelItem], error) {
	logger.Info("Getting all travel items")
	return s.Repository.GetAll(queryParams)
}
