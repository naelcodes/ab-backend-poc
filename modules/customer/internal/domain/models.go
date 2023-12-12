package domain

import "neema.co.za/rest/utils/types"

type Customer struct {
	types.BaseEntity
	CustomerName    string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string
}
