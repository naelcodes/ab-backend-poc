package booking

import (
	. "neema.co.za/rest/modules/booking/internal/api"
	"neema.co.za/rest/utils/managers"
)

func GetModule(dependencyManager *managers.DependencyManager) *Module {
	api := BuildApi(dependencyManager)
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("", api.GetAllTravelItemsHandler)
}
