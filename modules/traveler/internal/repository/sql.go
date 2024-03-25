package repository

import "neema.co.za/rest/utils/models"

func (r *Repository) GetAll() ([]*models.Traveler, error) {
	travelers := make([]*models.Traveler, 0)
	_ = r.Find(&travelers)
	return travelers, nil
}
