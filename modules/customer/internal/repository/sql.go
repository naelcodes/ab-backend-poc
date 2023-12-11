package repository

import (
	"fmt"

	"neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
)

// func (r *Repository) GetByID(id int) ([]*Customer, error) {
// 	var customers = make([]*Customer, 0)
// 	r.Find(customers)
// 	return customers, nil
// }

func (r *Repository) GetAll() ([]*Customer, error) {

	logger.Info("Getting all customers")

	var customers = make([]*Customer, 0)
	err := r.Where("tag = ?", "3").Find(&customers)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting customers: %v", err))
		return nil, err
	}

	logger.Info(fmt.Sprintf("Found %v customers", len(customers)))
	return customers, nil
}
