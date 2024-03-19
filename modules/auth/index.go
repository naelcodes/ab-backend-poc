package auth

import (
	. "neema.co.za/rest/modules/auth/internal/api"
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
	api.Post("/signin", middlewares.PayloadValidator(new(payloads.AuthSignInPayload)), api.EmailSignInHandler)
	api.Post("/signup", middlewares.PayloadValidator(new(payloads.AuthSignUpPayload)), api.EmailSignUpHandler)
	api.Post("/verification/:code", api.CodeVerificationHandler)
	api.Get("/facebook", api.FacebookAuthHandler)
	api.Get("/facebook/redirect", api.FacebookAuthRedirectHandler)
	api.Get("/google", api.GoogleAuthHandler)
	api.Get("/google/redirect", api.GoogleAuthRedirectHandler)
	//api.Post("/add-password")
}
