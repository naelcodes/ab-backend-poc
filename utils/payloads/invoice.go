package payloads

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"neema.co.za/rest/utils/models"
)

type CreateInvoicePayload struct {
	models.Invoice
	TravelItemIds []TravelItemPayload `json:"travelItems"`
}

type TravelItemDetails struct {
	Id int `json:"id"`
}

func (t TravelItemDetails) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Id, validation.Required),
	)
}

func (c CreateInvoicePayload) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IdCustomer, validation.Required),
		validation.Field(&c.CreationDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&c.DueDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&c.TravelItemIds, validation.Required),
	)
}

type UpdateInvoicePayload struct {
	Id           int                  `json:"id"`
	IdCustomer   *int                 `json:"idCustomer,omitempty"`
	CreationDate *string              `json:"creationDate,omitempty"`
	DueDate      *string              `json:"dueDate,omitempty"`
	TravelItems  *[]TravelItemPayload `json:"travelItems,omitempty"`
}

func (u UpdateInvoicePayload) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.IdCustomer, validation.NilOrNotEmpty),
		validation.Field(&u.CreationDate, validation.NilOrNotEmpty),
		validation.Field(&u.DueDate, validation.NilOrNotEmpty),
		validation.Field(&u.TravelItems, validation.NilOrNotEmpty),
	)
}
