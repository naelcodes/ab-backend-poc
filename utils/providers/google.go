package providers

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"neema.co.za/rest/utils/logger"
)

type GoogleProvider struct {
	Oauth2 *oauth2.Config
}

func NewGoogleProvider() *GoogleProvider {
	logger.Info(fmt.Sprintf("GOOGLE_CLIENT_ID : %s", os.Getenv("GOOGLE_CLIENT_ID")))
	logger.Info(fmt.Sprintf("GOOGLE_CLIENT_SECRET : %s", os.Getenv("GOOGLE_CLIENT_SECRET")))
	logger.Info(fmt.Sprintf("GOOGLE_REDIRECT_URL : %s", os.Getenv("GOOGLE_REDIRECT_URL")))

	return &GoogleProvider{&oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}}
}
