package repository

import (
	"fmt"

	"neema.co.za/rest/utils/dto"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/mappers"
	dbModels "neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"
)

func (r *Repository) Count() (*int64, error) {
	logger.Info("Counting travel items")

	totalRowCount, err := r.SQL("SELECT id FROM air_booking WHERE transaction_type = 'sales' AND product_type = 'flight' AND status = 'pending' AND id_invoice IS NULL").Count(new(dbModels.TravelItem))

	if err != nil {
		logger.Error(fmt.Sprintf("Error counting travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting travel items: %v", err))
	}

	logger.Info(fmt.Sprintf("Total travel item count: %v", totalRowCount))

	return &totalRowCount, nil
}

func (r *Repository) GetAll(queryParams *types.GetQueryParams) (*dto.GetAllTravelItemDTO, error) {
	logger.Info("Getting travel items")

	totalRowCount, err := r.Count()

	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := int(*totalRowCount)

	if queryParams != nil && queryParams.PageNumber != nil && queryParams.PageSize != nil {
		pageNumber = *queryParams.PageNumber
		pageSize = *queryParams.PageSize
	}

	logger.Info(fmt.Sprintf("PageNumber: %v", pageNumber))
	logger.Info(fmt.Sprintf("PageSize: %v", pageSize))

	var travelItems []*dbModels.TravelItem

	err = r.SQL("SELECT id,itinerary,traveler_name,ticket_number,CAST(total_price AS numeric) AS total_price FROM air_booking WHERE transaction_type = ? AND product_type = ? AND status = ? AND id_invoice IS NULL LIMIT ? OFFSET ?", "sales", "flight", "pending", pageSize, pageNumber*pageSize).Find(&travelItems)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting travel items: %v", err))
	}

	logger.Info(fmt.Sprintf("Total travel item count: %v", len(travelItems)))

	travelItemDTOList := TravelItemModelListToDTOList(travelItems)

	return &dto.GetAllTravelItemDTO{
		Data:          travelItemDTOList,
		TotalRowCount: int(*totalRowCount),
		PageSize:      pageSize,
		PageNumber:    pageNumber,
	}, nil
}
