// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package customer

import (
	"neema.co.za/rest/modules/customer/internal/api"
	"neema.co.za/rest/modules/customer/internal/repository"
	"neema.co.za/rest/modules/customer/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
)

// Injectors from wire.go:

// New api handler
func BuildApi() *api.Api {
	databaseDatabase := database.GetDatabase()
	repositoryRepository := repository.NewRepository(databaseDatabase)
	serviceService := service.NewService(repositoryRepository)
	fiberApp := app.NewFiberApp()
	apiApi := api.NewApi(serviceService, fiberApp)
	return apiApi
}
