package service

import (
	"neema.co.za/rest/modules/auth/internal/repository"
	"neema.co.za/rest/utils/managers"
	"neema.co.za/rest/utils/providers"
	"neema.co.za/rest/utils/sdk"
)

type Service struct {
	*repository.Repository
	*Imports
	*sdk.CasdoorSDK
	*providers.FacebookProvider
	*providers.GoogleProvider
}

func NewService(repository *repository.Repository, dependencyManager *managers.DependencyManager, sdk *sdk.CasdoorSDK, facebookProvider *providers.FacebookProvider, googleProvider *providers.GoogleProvider) *Service {
	return &Service{repository, &Imports{dependencyManager}, sdk, facebookProvider, googleProvider}
}
