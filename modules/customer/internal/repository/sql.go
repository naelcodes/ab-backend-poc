package repository

import (
	"fmt"

	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"

	CustomErrors "neema.co.za/rest/utils/errors"
)

func (r *Repository) Count() (int64, error) {
	logger.Info("Counting customers")

	customers := make([]*models.Customer, 0)
	totalRowCount, err := r.FindAndCount(&customers)

	if err != nil {
		logger.Error(fmt.Sprintf("Error counting customers: %v", err))
		return 0, CustomErrors.RepositoryError(fmt.Errorf("error counting customer records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total customer count: %v", totalRowCount))

	return totalRowCount, nil

}
func (r *Repository) GetAll(queryParams *types.GetQueryParams) (*types.GetAllDTO[[]*models.Customer], error) {

	customers := make([]*models.Customer, 0)

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

	err = r.Limit(pageSize, pageNumber*pageSize).Find(&customers)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting customers: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer records: %v", err))
	}

	logger.Info(fmt.Sprintf("Found %v customers", len(customers)))

	getAllCustomersDTO := new(types.GetAllDTO[[]*models.Customer])
	getAllCustomersDTO.Data = customers
	getAllCustomersDTO.TotalRowCount = int(totalRowCount)
	getAllCustomersDTO.PageNumber = pageNumber
	getAllCustomersDTO.PageSize = pageSize

	return getAllCustomersDTO, nil
}

func (r *Repository) GetById(id int) (*models.Customer, error) {

	customer := new(models.Customer)
	has, err := r.ID(id).Get(customer)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting customer: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer record: %v", err))
	}

	if !has {
		logger.Error(fmt.Sprintf("Customer with id %v not found", id))
		return nil, CustomErrors.NotFoundError(fmt.Errorf("customer with id %v not found", id))
	}

	logger.Info(fmt.Sprintf("Found customer: %v", customer))

	return customer, nil
}

func (r *Repository) Save(customer *models.Customer) (*models.Customer, error) {

	_, err := r.Insert(customer)

	if err != nil {
		logger.Error(fmt.Sprintf("Error saving customer: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving customer record: %v", err))
	}

	logger.Info(fmt.Sprintf("Saved customer: %v", customer))
	return r.GetById(customer.Id)
}

func (r *Repository) Update(id int, customer *models.Customer) error {

	has, err := r.Exist(&models.Customer{
		Id: id,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Error checking if customer exists: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error checking if customer exists: %v", err))
	}

	if !has {
		logger.Error(fmt.Sprintf("Customer with id %v not found", id))
		return CustomErrors.NotFoundError(fmt.Errorf("customer with id %v not found", id))
	}

	_, err = r.ID(id).Update(customer)

	if err != nil {
		logger.Error(fmt.Sprintf("Error updating customer: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error updating customer record: %v", err))
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	logger.Info(fmt.Sprintf("Deleting customer: %v", id))

	has, err := r.Exist(&models.Customer{
		Id: id,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Error checking if customer exists: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error checking if customer exists: %v", err))
	}

	if !has {
		logger.Error(fmt.Sprintf("Customer with id %v not found", id))
		return CustomErrors.NotFoundError(fmt.Errorf("customer with id %v not found", id))
	}

	_, err = r.ID(id).Delete(&models.Customer{})

	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting customer: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error deleting customer record: %v", err))
	}
	return nil
}
