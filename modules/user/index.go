package user

import (
	. "neema.co.za/rest/modules/user/api"
	. "neema.co.za/rest/utils/app"
)

const API_BASE_PATH = "/api/v1"

func Start(handler *Api) {

	router := handler.RouterCreator(API_BASE_PATH)
	handleRoutes(handler, router.Group("/customers"))
}

func handleRoutes(handler *Api, router Router) {
	router.Get("", handler.GetUserByID)
}
