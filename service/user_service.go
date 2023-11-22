// service/user_service.go
package service

import (
	"neema.co.za/rest/models"
	"neema.co.za/rest/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	// Business logic (if any)
	return s.userRepository.GetUserByID(id)
}
