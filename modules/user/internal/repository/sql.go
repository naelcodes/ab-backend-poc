package repository

import (
	. "neema.co.za/rest/utils/models"
)

func (r *Repository) GetUserByID(id int) (*User, error) {
	var users []User
	r.Find(users)
	return &users[0], nil
}
