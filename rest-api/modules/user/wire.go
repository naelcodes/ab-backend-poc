//go:build wireinject

package user

import (
	"github.com/google/wire"
	"neema.co.za/rest/modules/user/internal/api"
	"neema.co.za/rest/modules/user/internal/repository"
	"neema.co.za/rest/modules/user/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/managers"
)

func BuildApi(dependencyManager *managers.DependencyManager) *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
