package repository

import (
	"neema.co.za/rest/modules/customer/internal/domain"
	"neema.co.za/rest/modules/customer/internal/dto"
	"neema.co.za/rest/utils/models"
)

func CustomerModelToDTO(customer *models.Customer) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = customer.Id
	customerDTO.Customer_name = customer.Customer_name
	customerDTO.State = customer.State
	customerDTO.Account_number = customer.Account_number
	customerDTO.Alias = customer.Alias
	customerDTO.Ab_key = customer.Ab_key
	customerDTO.Tmc_client_number = customer.Tmc_client_number

	return customerDTO
}

func CustomerModelListToDTOList(customers []*models.Customer) []*dto.GetCustomerDTO {
	customerDTOList := make([]*dto.GetCustomerDTO, 0)

	for _, customer := range customers {
		customerDTOList = append(customerDTOList, CustomerModelToDTO(customer))
	}

	return customerDTOList
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
