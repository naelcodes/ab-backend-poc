package traveler

import (
	. "neema.co.za/rest/modules/traveler/internal/api"
	"neema.co.za/rest/utils/managers"
)

func GetModule(dependencyManager *managers.DependencyManager) *Module {
	api := BuildApi(dependencyManager)
	handleRoutes(api)
	module := Module(*api)                //Module is an alias of Api
	dependencyManager.Add(module.Exports) //add exportable functions to dependency manager
	return &module
}

func handleRoutes(api *Api) {
	api.Get("", api.GetAllTravelersHandler)
	api.Post("", api.CreateTravelerHandler)
	api.Get("/:id", api.GetSingleTravelerHandler)
	api.Put("/:id", api.UpdateTravelerHandler)
	api.Delete("/:id", api.DeleteTravelerHandler)
}
