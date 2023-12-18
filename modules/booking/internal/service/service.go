package service

import "neema.co.za/rest/modules/booking/internal/repository"

type Service struct {
	*repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository}
}
