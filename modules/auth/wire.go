//go:build wireinject

package auth

import (
	"github.com/google/wire"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/managers"
	"neema.co.za/rest/utils/providers"
	"neema.co.za/rest/utils/sdk"

	"neema.co.za/rest/modules/auth/internal/api"
	"neema.co.za/rest/modules/auth/internal/repository"
	"neema.co.za/rest/modules/auth/internal/service"
)

func BuildApi(dependencyManager *managers.DependencyManager) *api.Api {
	panic(wire.Build(
		providers.NewFacebookProvider,
		providers.NewGoogleProvider,
		sdk.NewCasdoorSdk,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
