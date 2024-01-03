package customer

import (
	. "neema.co.za/rest/modules/customer/internal/api"
	"neema.co.za/rest/utils/middlewares"
	"neema.co.za/rest/utils/payloads"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("", api.GetAllCustomerHandler)
	api.Get("/:id", api.GetCustomerHandler)
	api.Post("", middlewares.PayloadValidator(new(payloads.CreateCustomerPayload)), api.CreateCustomerHandler)

}
