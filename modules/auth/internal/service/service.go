package service

import (
	"neema.co.za/rest/modules/auth/internal/repository"
	"neema.co.za/rest/modules/auth/internal/utils/providers"
	"neema.co.za/rest/modules/auth/internal/utils/sdk"
	"neema.co.za/rest/modules/auth/internal/utils/sessions"
	"neema.co.za/rest/utils/managers"
)

type Service struct {
	*repository.Repository
	*Imports
	*sdk.CasdoorSDK
	*providers.FacebookProvider
	*providers.GoogleProvider
	*sessions.AppSessionStore
}

func NewService(repository *repository.Repository, dependencyManager *managers.DependencyManager, sdk *sdk.CasdoorSDK, facebookProvider *providers.FacebookProvider, googleProvider *providers.GoogleProvider, appSessionStore *sessions.AppSessionStore) *Service {
	return &Service{repository, &Imports{dependencyManager}, sdk, facebookProvider, googleProvider, appSessionStore}
}
