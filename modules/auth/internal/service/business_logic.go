package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/huandu/facebook/v2"
	"golang.org/x/crypto/bcrypt"
	Mailer "neema.co.za/rest/modules/auth/internal/utils/mail"
	"neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/models"

	CustomErrors "neema.co.za/rest/utils/errors"
)

func (s *Service) EmailSignInService(user *models.User) (bool, error) {

	logger.Info(fmt.Sprintf("Email: %s", user.Email))
	savedUser, _ := s.CasdoorSDK.GetUserByEmail(user.Email)
	if savedUser != nil {
		logger.Info(fmt.Sprintf("Found User %v", savedUser.Name))
		err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(user.Password))
		if err != nil {
			logger.Error(fmt.Sprintf("Password Mismatch %s", err.Error()))
			return false, CustomErrors.NotFoundError(fmt.Errorf("user with credentials not found"))
		}

		newSessionID, _ := s.GenerateSessionID()
		err = s.Storage.Set(newSessionID, []byte(savedUser.Id), 30*time.Minute)

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to set session %s", err.Error()))
			return false, CustomErrors.ServiceError(err, "Failed to set session")
		}
		return true, nil
	}

	return false, CustomErrors.NotFoundError(fmt.Errorf("user with email %s not found", user.Email))

}

func (s *Service) BeginEmailSignUpService(user *models.User) (bool, error) {

	savedUser, _ := s.CasdoorSDK.GetUserByEmail(user.Email)

	if savedUser != nil {
		logger.Info(fmt.Sprintf("Found User %v", savedUser.Name))
		return false, CustomErrors.DomainError(fmt.Errorf("user with email %s already exists", user.Email))
	}

	verificationCode := helpers.GenerateRandomCode()
	jsonData, _ := json.Marshal(user)

	// store the user data temporarily in redis
	err := s.Repository.RedisStore.Set(context.Background(), fmt.Sprintf("%v", verificationCode), string(jsonData), time.Minute*30).Err()

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to set verification code %s", err.Error()))
		return false, CustomErrors.ServiceError(err, "Failed to set verification code")
	}

	err = Mailer.SendMail(user.Email, "Email Verification", fmt.Sprintf("Your verification code is: %v", verificationCode))

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to send verification email %s", err.Error()))
		return false, CustomErrors.ServiceError(err, "Failed to send verification mail")
	}
	return true, nil
}

func (s *Service) CompleteEmailSignUpService(code int) (*models.User, error) {

	result, err := s.Repository.RedisStore.Get(context.Background(), fmt.Sprintf("%v", code)).Result()

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get verification code %s", err.Error()))
		return nil, CustomErrors.NotFoundError(fmt.Errorf("verification code expired"))
	}
	user := models.User{}
	err = json.Unmarshal([]byte(result), &user)

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to convert json %s", err.Error()))
	}

	casdoorUser := new(casdoorsdk.User)
	casdoorUser.Email = user.Email
	casdoorUser.Name = user.Username

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(fmt.Sprintf("Password Hashing Failed: %s", err.Error()))
		return nil, CustomErrors.ServiceError(err, "Password Hashing Failed")
	}

	casdoorUser.Password = string(hash)
	casdoorUser.EmailVerified = true

	_, err = s.CasdoorSDK.AddUser(casdoorUser)

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to add user %s", err.Error()))
		return nil, CustomErrors.ServiceError(err, "Failed to add user")
	}

	s.Storage.Delete(fmt.Sprintf("%v", code))

	casdoorUser, _ = s.CasdoorSDK.GetUserByEmail(user.Email)
	user.Id = casdoorUser.Id
	return &user, nil

}

func (s *Service) FacebookAuthService() string {
	// TODO : state should be a random string value to protect against CSRF attacks
	url := s.FacebookProvider.Oauth2.AuthCodeURL("state")
	return url
}

func (s *Service) FacebookAuthRedirectService(code string, state string) (bool, *models.User, error) {
	//TODO: Validate state value received from facebook

	token, err := s.FacebookProvider.Oauth2.Exchange(context.Background(), code)

	if err != nil {
		logger.Error(fmt.Sprintf("Code-Token Exchange failed: %s", err.Error()))
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("Code-Token Exchange failed: %w", err), "Facebook Code-Token Exchange")
	}

	result, err := s.FacebookProvider.GraphAPI.Session(token.AccessToken).Get("/me", facebook.Params{
		"fields": "id,email,first_name,last_name",
	})

	if err != nil {
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("failed to fetch facebook user: %w", err), "facebook graph api fetch")
	}

	user := models.User{
		Email:    result.Get("email").(string),
		Password: "",
		Username: result.Get("first_name").(string) + " " + result.Get("last_name").(string),
	}

	return s.SocialAuthService(&user)
}

func (s *Service) GoogleAuthService() string {
	// TODO : state should be a random string value to protect against CSRF attacks
	url := s.GoogleProvider.Oauth2.AuthCodeURL("state")
	return url
}

func (s *Service) GoogleAuthRedirectService(code string, state string) (bool, *models.User, error) {
	//TODO: Validate state value received from google

	token, err := s.GoogleProvider.Oauth2.Exchange(context.Background(), code)
	if err != nil {
		logger.Error(fmt.Sprintf("Code-Token Exchange failed: %s", err.Error()))
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("Code-Token Exchange failed: %w", err), "Google Code-Token Exchange")
	}

	response, err := http.Get(fmt.Sprintf("%v?access_token=%v", os.Getenv("GOOGLE_OAUTH2_USERINFO_URL"), token.AccessToken))

	if err != nil {
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("failed to fetch google user data: %w", err), "Google graph api fetch")
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("JSON parsing Failed: %w", err), "Parsing google user info in Response Body")
	}

	GoogleUser := models.GoogleUser{}
	err = json.Unmarshal(data, &GoogleUser)
	if err != nil {
		return false, nil, CustomErrors.ServiceError(fmt.Errorf("JSON parsing Failed: %w", err), "Json Unmarshal ")
	}

	user := models.User{
		Email:    GoogleUser.Email,
		Password: "",
		Username: GoogleUser.FirstName + " " + GoogleUser.LastName,
	}
	return s.SocialAuthService(&user)
}

func (s *Service) SocialAuthService(user *models.User) (isNew bool, newUser *models.User, err error) {

	savedUser, _ := s.CasdoorSDK.GetUserByEmail(user.Email)

	if savedUser != nil {
		isNew = false

		logger.Info(fmt.Sprintf("User %s already exists", user.Email))

		newSessionID, _ := s.GenerateSessionID()
		err := s.Storage.Set(newSessionID, []byte(savedUser.Id), 30*time.Minute)

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to set session %s", err.Error()))
			return false, nil, CustomErrors.ServiceError(err, "Failed to set session")
		}

		return isNew, nil, nil

	} else {

		isNew = true

		logger.Info(fmt.Sprintf("Creating user %s", user.Email))

		casdoorUser := new(casdoorsdk.User)
		casdoorUser.Email = user.Email
		casdoorUser.Name = user.Username
		casdoorUser.EmailVerified = true

		_, err := s.CasdoorSDK.AddUser(casdoorUser)

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to create user %s", err.Error()))
			return false, nil, CustomErrors.ServiceError(err, "Failed to create user")
		}

		casdoorUser, _ = s.CasdoorSDK.GetUserByEmail(user.Email)
		user.Id = casdoorUser.Id

		return isNew, user, nil
	}

}

func (s *Service) AddPasswordService(user *models.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(fmt.Sprintf("Password Hashing Failed: %s", err.Error()))
		return CustomErrors.ServiceError(err, "Password Hashing Failed")
	}

	savedUser, _ := s.CasdoorSDK.GetUserByEmail(user.Email)

	if savedUser != nil {
		savedUser.Password = string(hash)
		_, err := s.CasdoorSDK.UpdateUser(savedUser)

		if err != nil {
			logger.Error(fmt.Sprintf("Password Update Failed: %s", err.Error()))
			return CustomErrors.ServiceError(err, "Password Update Failed")
		}

		return nil
	}

	return CustomErrors.NotFoundError(fmt.Errorf("user with email %s not found", user.Email))

}
