//go:build wireinject

package imputation

import (
	"github.com/google/wire"
	"neema.co.za/rest/modules/imputation/internal/api"
	"neema.co.za/rest/modules/imputation/internal/repository"
	"neema.co.za/rest/modules/imputation/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
)

func BuildApi() *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
