package traveler

import (
	. "neema.co.za/rest/modules/traveler/internal/api"
	"neema.co.za/rest/utils/managers"
	"neema.co.za/rest/utils/middlewares"
	"neema.co.za/rest/utils/payloads"
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
	api.Post("", middlewares.PayloadValidator(new(payloads.CreateTravelerPayload)), api.CreateTravelerHandler)
	api.Get("/:id", api.GetTravelerByIdHandler)
	api.Put("/:id", middlewares.PayloadValidator(new(payloads.UpdateTravelerPayload)), api.UpdateTravelerHandler)
	// api.Delete("/:id", api.DeleteTravelerHandler)
}
