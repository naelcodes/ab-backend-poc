package user

import (
	"github.com/gofiber/fiber/v2"

	. "neema.co.za/rest/modules/user/internal/api"
)

func GetApp() *fiber.App {
	api := GetApi()
	handleRoutes(api)
	return api.App
}

func handleRoutes(api *Api) {
	api.Get("", api.GetUserByID)
}
