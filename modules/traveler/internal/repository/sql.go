package repository

import (
	"fmt"

	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"

	CustomErrors "neema.co.za/rest/utils/errors"
)

func (r *Repository) Count() (int64, error) {
	logger.Info("Counting travelers")
	travelers := make([]*models.Traveler, 0)
	totalRowCount, err := r.FindAndCount(&travelers)

	if err != nil {
		logger.Error(fmt.Sprintf("Error counting travelers: %v", err))
		return 0, CustomErrors.RepositoryError(fmt.Errorf("error counting traveler records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total traveler count: %v", totalRowCount))
	return totalRowCount, nil
}

func (r *Repository) GetAll(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.Traveler], error) {
	travelers := make([]*models.Traveler, 0)

	totalRowCount, err := r.Count()

	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := int(totalRowCount)

	logger.Info(fmt.Sprintf("QueryParams: %v", queryParams))

	if queryParams != nil {

		if queryParams.PageNumber != nil && queryParams.PageSize != nil {
			pageNumber = *queryParams.PageNumber
			pageSize = *queryParams.PageSize
		}

		logger.Info(fmt.Sprintf("PageNumber: %v", pageNumber))
		logger.Info(fmt.Sprintf("PageSize: %v", pageSize))
	}

	err = r.Limit(pageSize, pageNumber*pageSize).Find(&travelers)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting travelers: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting traveler records: %v", err))
	}

	logger.Info(fmt.Sprintf("Found %v travelers", len(travelers)))

	result := types.GetAllDTO[[]*models.Traveler]{
		Data:          travelers,
		TotalRowCount: int(totalRowCount),
		PageSize:      pageSize,
		PageNumber:    pageNumber,
	}

	return &result, nil
}
