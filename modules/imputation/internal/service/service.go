package service

import "neema.co.za/rest/modules/imputation/internal/repository"

type Service struct {
	*repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository}
}
