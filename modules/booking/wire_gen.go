// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package booking

import (
	"neema.co.za/rest/modules/booking/internal/api"
	"neema.co.za/rest/modules/booking/internal/repository"
	"neema.co.za/rest/modules/booking/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
)

// Injectors from wire.go:

func BuildApi() *api.Api {
	databaseDatabase := database.GetDatabase()
	repositoryRepository := repository.NewRepository(databaseDatabase)
	serviceService := service.NewService(repositoryRepository)
	fiberApp := app.NewFiberApp()
	apiApi := api.NewApi(serviceService, fiberApp)
	return apiApi
}
