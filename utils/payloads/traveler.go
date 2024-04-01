package payloads

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"neema.co.za/rest/utils/models"
)

type CreateTravelerPayload struct {
	models.Traveler
}

func (c CreateTravelerPayload) Validate() error {
	return validation.ValidateStruct(&c, validation.Field(&c.DisplayName, validation.Required))
}
