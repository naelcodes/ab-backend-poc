// handlers/api.go
package api

import (
	"neema.co.za/rest/modules/user/service"
	"neema.co.za/rest/utils/app"
)

type Api struct {
	*service.Service
	app.RouterCreator
}

func NewApi(service *service.Service, routerCreator app.RouterCreator) *Api {
	return &Api{service, routerCreator}
}
