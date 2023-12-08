// handlers/api.go
package api

import (
	"github.com/gofiber/fiber/v2"

	"neema.co.za/rest/modules/user/internal/service"
)

type Module Api
type Api struct {
	*service.Service
	*fiber.App
}

func NewApi(service *service.Service, app *fiber.App) *Api {
	return &Api{service, app}
}
