//go:build wireinject

package main

import (
	"github.com/google/wire"

	"neema.co.za/rest/api/handlers"
	"neema.co.za/rest/app"
	"neema.co.za/rest/database"
	"neema.co.za/rest/repository"
	"neema.co.za/rest/service"
)

func injectUserModule() *handlers.UserHandler {
	panic(wire.Build(
		database.GetDatabase,
		repository.NewUserRepository,
		service.NewUserService,
		handlers.NewUserHandler,
		app.NewRouter,
	))
}
