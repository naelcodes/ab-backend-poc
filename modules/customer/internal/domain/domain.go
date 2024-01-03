package domain

import (
	"neema.co.za/rest/utils/models"
)

type CustomerDomain struct {
	customer *models.Customer
	builder  *CustomerBuilder
}

func NewCustomerDomain() *CustomerDomain {
	domain := &CustomerDomain{customer: &models.Customer{}}
	return domain
}

func (domain *CustomerDomain) GetCustomerBuilder() *CustomerBuilder {
	if domain.builder == nil {
		domain.builder = NewCustomerBuilder(domain.customer)
	}
	return domain.builder
}

func (domain *CustomerDomain) GetCustomer() *models.Customer {
	return domain.customer
}
