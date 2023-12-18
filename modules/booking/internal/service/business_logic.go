package service

import (
	"neema.co.za/rest/utils/dto"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllTravelItemsService(queryParams *types.GetQueryParams) (*dto.GetAllTravelItemDTO, error) {
	logger.Info("Getting all travel items")
	return s.Repository.GetAll(queryParams)
}
