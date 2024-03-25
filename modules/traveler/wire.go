//go:build wireinject

package traveler

import (
	"github.com/google/wire"

	"neema.co.za/rest/modules/traveler/internal/api"
	"neema.co.za/rest/modules/traveler/internal/repository"
	"neema.co.za/rest/modules/traveler/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/managers"
)

// New api handler
func BuildApi(dependencyManager *managers.DependencyManager) *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
