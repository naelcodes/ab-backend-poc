package imputation

import (
	. "neema.co.za/rest/modules/imputation/internal/api"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {

}
