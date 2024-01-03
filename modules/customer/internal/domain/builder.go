package domain

import (
	"neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/models"
)

type CustomerBuilder struct {
	customer *models.Customer
}

func NewCustomerBuilder(customer *models.Customer) *CustomerBuilder {
	builder := new(CustomerBuilder)
	builder.customer = customer
	return builder
}

func (builder *CustomerBuilder) SetId(id int) *CustomerBuilder {
	builder.customer.Id = id
	return builder
}

func (builder *CustomerBuilder) SetCustomerName(name string) *CustomerBuilder {
	builder.customer.CustomerName = name
	return builder
}

func (builder *CustomerBuilder) SetAlias(alias string) *CustomerBuilder {
	builder.customer.Alias = alias
	return builder
}

func (builder *CustomerBuilder) SetAbKey() *CustomerBuilder {
	builder.customer.AbKey = helpers.GenerateRandomString(15)
	return builder
}

func (builder *CustomerBuilder) SetTmcClientNumber(tmcClientNumber string) *CustomerBuilder {
	builder.customer.TmcClientNumber = tmcClientNumber
	return builder
}

func (builder *CustomerBuilder) SetAccountNumber(accountNumber string) *CustomerBuilder {
	builder.customer.AccountNumber = accountNumber
	return builder
}

func (builder *CustomerBuilder) SetState(state string) *CustomerBuilder {
	builder.customer.State = state
	return builder
}

func (builder *CustomerBuilder) SetDefaults() *CustomerBuilder {
	builder.customer.Tag = "3"
	builder.customer.IdCountry = 40
	builder.customer.IdCurrency = 550
	return builder
}
