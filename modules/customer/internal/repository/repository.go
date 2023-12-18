// repository/repository.go
package repository

import (
	"neema.co.za/rest/modules/customer/internal/domain"
	. "neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/models"
)

type Repository struct {
	*Database
}

func NewRepository(database *Database) *Repository {
	return &Repository{database}
}

func CustomerModelFromDomainModel(customerDomainModel *domain.Customer) *models.Customer {
	customerModel := new(models.Customer)

	customerModel.Id = int(customerDomainModel.Id)
	customerModel.Customer_name = customerDomainModel.CustomerName
	customerModel.State = customerDomainModel.State
	customerModel.Account_number = customerDomainModel.AccountNumber
	customerModel.Alias = customerDomainModel.Alias
	customerModel.Ab_key = customerDomainModel.AbKey
	customerModel.Tmc_client_number = customerDomainModel.TmcClientNumber
	customerModel.Tag = "3"
	customerModel.Id_country = 40
	customerModel.Id_currency = 550
	return customerModel

}
