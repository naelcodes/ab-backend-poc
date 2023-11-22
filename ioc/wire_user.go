//go:build wireinject

package ioc

import (
	"github.com/google/wire"

	"neema.co.za/rest/modules/user/api"
	"neema.co.za/rest/modules/user/repository"
	"neema.co.za/rest/modules/user/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
)

func InjectUserModule() *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
		app.NewRouter,
	))
}
