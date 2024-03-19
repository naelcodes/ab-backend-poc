package payloads

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"neema.co.za/rest/utils/models"
)

type AuthSignInPayload struct {
	models.User
}

func (a AuthSignInPayload) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}

type AuthSignUpPayload struct {
	models.User
}

func (a AuthSignUpPayload) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password, validation.Required),
		validation.Field(&a.Username, validation.Required),
	)
}
