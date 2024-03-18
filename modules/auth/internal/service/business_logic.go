package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/huandu/facebook/v2"
	"golang.org/x/crypto/bcrypt"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"

	CustomErrors "neema.co.za/rest/utils/errors"
)

func (s *Service) EmailSignInService(email string, password string) {

	savedUser, _ := s.CasdoorSDK.GetUserByEmail(email)
	if savedUser != nil {
		logger.Info(fmt.Sprintf("Found User %v", savedUser.Name))
		err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(password))
		if err != nil {
			logger.Error(fmt.Sprintf("Password Mismatch %s", err.Error()))
			return //sign in failed
		}
		return //sign in success
	}

}

func (s *Service) EmailSignUpService(payload models.User) {
	user := new(casdoorsdk.User)

	user.Name = payload.Username
	user.Email = payload.Email

}

func (s *Service) FacebookAuthService() string {
	// TODO : state should be a random string value to protect against CSRF attacks
	url := s.FacebookProvider.Oauth2.AuthCodeURL("state")
	return url
}

func (s *Service) FacebookAuthRedirectService(code string, state string) (*models.FacebookUser, error) {
	//TODO: Validate state value received from facebook

	token, err := s.FacebookProvider.Oauth2.Exchange(context.Background(), code)

	if err != nil {
		logger.Error(fmt.Sprintf("Code-Token Exchange failed: %s", err.Error()))
		return nil, CustomErrors.ServiceError(fmt.Errorf("Code-Token Exchange failed: %w", err), "Facebook Code-Token Exchange")
	}

	result, err := s.FacebookProvider.GraphAPI.Session(token.AccessToken).Get("/me", facebook.Params{
		"fields": "id,email,first_name,last_name",
	})

	if err != nil {
		return nil, CustomErrors.ServiceError(fmt.Errorf("failed to fetch facebook user: %w", err), "facebook graph api fetch")
	}
	user := models.FacebookUser{
		Id:        result.Get("id").(string),
		Email:     result.Get("email").(string),
		FirstName: result.Get("first_name").(string),
		LastName:  result.Get("last_name").(string),
	}

	return &user, nil
}

func (s *Service) GoogleAuthService() string {
	// TODO : state should be a random string value to protect against CSRF attacks
	url := s.GoogleProvider.Oauth2.AuthCodeURL("state")
	return url
}

func (s *Service) GoogleAuthRedirectService(code string, state string) (*models.GoogleUser, error) {
	//TODO: Validate state value received from google

	token, err := s.GoogleProvider.Oauth2.Exchange(context.Background(), code)
	if err != nil {
		logger.Error(fmt.Sprintf("Code-Token Exchange failed: %s", err.Error()))
		return nil, CustomErrors.ServiceError(fmt.Errorf("Code-Token Exchange failed: %w", err), "Google Code-Token Exchange")
	}

	response, err := http.Get(fmt.Sprintf("%v?access_token=%v", os.Getenv("GOOGLE_OAUTH2_USERINFO_URL"), token.AccessToken))

	if err != nil {
		return nil, CustomErrors.ServiceError(fmt.Errorf("failed to fetch google user data: %w", err), "Google graph api fetch")
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, CustomErrors.ServiceError(fmt.Errorf("JSON parsing Failed: %w", err), "Parsing google user info in Response Body")
	}

	user := models.GoogleUser{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, CustomErrors.ServiceError(fmt.Errorf("JSON parsing Failed: %w", err), "Json Unmarshal ")
	}
	return &user, nil
}

func (s *Service) SocialAuthService(email string) {

	savedUser, _ := s.CasdoorSDK.GetUserByEmail(email)

	if savedUser != nil {
		// TODO: Update User Session
		logger.Info(fmt.Sprintf("User %s already exists", email))
		return
	} else {

		logger.Info(fmt.Sprintf("Creating user %s", email))
		return
	}

}

func (s *Service) AddPasswordService(userId string, email string, password string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(fmt.Sprintf("Password Hashing Failed: %s", err.Error()))
		return
	}

	savedUser, _ := s.CasdoorSDK.GetUserByUserId(userId)

	if savedUser != nil {
		savedUser.Password = string(hash)
		_, err := s.CasdoorSDK.UpdateUser(savedUser)

		if err != nil {
			logger.Error(fmt.Sprintf("Password Update Failed: %s", err.Error()))
			return
		}

		logger.Info(fmt.Sprintf("User %s password updated", email))
		return
	}

}
