package customer

import (
	. "neema.co.za/rest/modules/customer/internal/api"
	"neema.co.za/rest/utils/dto"
	"neema.co.za/rest/utils/middlewares"
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
	api.Post("", middlewares.PayloadValidator(new(dto.CreateCustomerDTO)), api.CreateCustomerHandler)

}
