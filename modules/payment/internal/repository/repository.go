package repository

import (
	"neema.co.za/rest/modules/payment/internal/domain"
	. "neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/models"
)

type Repository struct {
	*Database
}

func NewRepository(database *Database) *Repository {
	return &Repository{database}
}

func PaymentModelFromDomainModel(paymentDomainModel *domain.Payment, paymentRecordCount int) *models.Payment {

	paymentModel := new(models.Payment)
	paymentModel.Id = int(paymentDomainModel.Id)
	paymentModel.Payment_number = helpers.GenerateCode("pr", paymentRecordCount+1)
	paymentModel.Payment_date = paymentDomainModel.PaymentDate
	paymentModel.Fop = paymentDomainModel.PaymentMode
	paymentModel.Amount = paymentDomainModel.Amount
	paymentModel.Balance = paymentDomainModel.Balance
	paymentModel.Used_amount = paymentDomainModel.UsedAmount
	paymentModel.Status = paymentDomainModel.Status
	paymentModel.Id_customer = &paymentDomainModel.IdCustomer
	paymentModel.Base_amount = paymentDomainModel.Amount
	paymentModel.Type = "customer_payment"
	paymentModel.Id_currency = 550
	paymentModel.Id_chart_of_accounts = 39
	paymentModel.Tag = "3"
	return paymentModel
}
