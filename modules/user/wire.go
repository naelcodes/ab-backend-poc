//go:build wireinject

package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"

	"neema.co.za/rest/modules/user/internal/api"
	"neema.co.za/rest/modules/user/internal/repository"
	"neema.co.za/rest/modules/user/internal/service"
	"neema.co.za/rest/utils/database"
)

// New api handler
func GetApi() *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}

func NewFiberApp() *fiber.App {
	return fiber.New()
}
