//go:build wireinject

package auth

import (
	"github.com/google/wire"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/managers"

	"neema.co.za/rest/modules/auth/internal/api"
	"neema.co.za/rest/modules/auth/internal/repository"
	"neema.co.za/rest/modules/auth/internal/service"
	"neema.co.za/rest/modules/auth/internal/utils/providers"
	"neema.co.za/rest/modules/auth/internal/utils/sdk"
	"neema.co.za/rest/modules/auth/internal/utils/sessions"
)

func BuildApi(dependencyManager *managers.DependencyManager) *api.Api {
	panic(wire.Build(
		sessions.NewAppSessionStore,
		database.GetRedisStore,
		providers.NewFacebookProvider,
		providers.NewGoogleProvider,
		sdk.NewCasdoorSdk,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
