// service/service.go
package service

import (
	"neema.co.za/rest/modules/traveler/internal/repository"
	"neema.co.za/rest/utils/managers"
)

type Service struct {
	*repository.Repository
	*Imports
}

func NewService(repository *repository.Repository, dependencyManager *managers.DependencyManager) *Service {
	return &Service{repository, &Imports{dependencyManager}}
}
