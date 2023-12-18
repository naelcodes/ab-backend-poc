package booking

import (
	. "neema.co.za/rest/modules/booking/internal/api"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("", api.GetAllTravelItemsHandler)
}
