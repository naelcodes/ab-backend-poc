package providers

import (
	"fmt"
	"os"

	fb "github.com/huandu/facebook/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"neema.co.za/rest/utils/logger"
)

type FacebookProvider struct {
	Oauth2   *oauth2.Config
	GraphAPI *fb.App
}

func NewFacebookProvider() *FacebookProvider {

	logger.Info(fmt.Sprintf("FACEBOOK_CLIENT_ID : %s", os.Getenv("FACEBOOK_CLIENT_ID")))
	logger.Info(fmt.Sprintf("FACEBOOK_CLIENT_SECRET : %s", os.Getenv("FACEBOOK_CLIENT_SECRET")))
	logger.Info(fmt.Sprintf("FACEBOOK_REDIRECT_URL : %s", os.Getenv("FACEBOOK_REDIRECT_URL")))

	return &FacebookProvider{&oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}, &fb.App{
		AppId:     os.Getenv("FACEBOOK_CLIENT_ID"),
		AppSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	},
	}
}
