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

func (r *Repository) GetById(id int) (*models.Traveler, error) {
	logger.Info(fmt.Sprintf("Getting traveler: %v", id))

	traveler := new(models.Traveler)
	has, err := r.ID(id).Get(traveler)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting traveler: %v", id))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting traveler: %v", id))
	}

	if !has {
		logger.Error(fmt.Sprintf("Traveler with id %v not found", id))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("traveler with id %v not found", id))
	}

	logger.Info(fmt.Sprintf("Found traveler: %v", traveler))
	return traveler, nil
}

func (r *Repository) Save(traveler *models.Traveler) (*models.Traveler, error) {
	logger.Info(fmt.Sprintf("Saving traveler: %v", traveler))

	_, err := r.Insert(traveler)
	if err != nil {
		logger.Error(fmt.Sprintf("Error saving traveler: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving traveler: %v", err))
	}

	logger.Info(fmt.Sprintf("Saved traveler: %v", traveler))
	return r.GetById(traveler.Id)
}

func (r *Repository) Update(id int, traveler *models.Traveler) (*models.Traveler, error) {
	logger.Info(fmt.Sprintf("Updating traveler: %v", traveler))

	has, err := r.Exist(&models.Traveler{Id: id})

	if err != nil {
		logger.Error(fmt.Sprintf("Error checking if customer exists: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error checking if customer exists : %v", err))
	}

	if !has {
		logger.Error(fmt.Sprintf("Traveler with id %v not found", id))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("traveler with id %v not found", id))
	}

	_, err = r.ID(id).Update(traveler)

	if err != nil {
		logger.Error(fmt.Sprintf("Error updating traveler: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error updating traveler: %v", err))
	}
	return r.GetById(id)

}
