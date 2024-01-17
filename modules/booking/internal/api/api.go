package api

import (
	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/modules/booking/internal/service"
)

type Module Api
type Api struct {
	*service.Service
	*fiber.App
	*Exports
}

func NewApi(service *service.Service, app *fiber.App) *Api {
	return &Api{service, app, &Exports{internalService: service}}
}
