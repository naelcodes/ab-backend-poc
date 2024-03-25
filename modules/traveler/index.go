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
	api.Get("/travelers", api.GetAllTravelersHandler)
	api.Post("/travelers", api.CreateTravelerHandler)
	api.Get("/travelers/:id", api.GetSingleTravelerHandler)
	api.Put("/travelers/:id", api.UpdateTravelerHandler)
	api.Delete("/travelers/:id", api.DeleteTravelerHandler)
}
