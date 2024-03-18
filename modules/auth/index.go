package auth

import (
	. "neema.co.za/rest/modules/auth/internal/api"
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
	api.Post("", api.EmailSignInHandler)
	api.Post("/signup", api.EmailSignUpHandler)
	api.Post("/email-verification", api.EmailVerificationHandler)
	api.Get("/facebook", api.FacebookAuthHandler)
	api.Get("/facebook/redirect", api.FacebookAuthRedirectHandler)
	api.Get("/google", api.GoogleAuthHandler)
	api.Get("/google/redirect", api.GoogleAuthRedirectHandler)
}
