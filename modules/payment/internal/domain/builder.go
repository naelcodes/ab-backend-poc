package domain

import (
	Helpers "neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/models"
)

type PaymentBuilder struct {
	payment *models.Payment
}

func NewPaymentBuilder(payment *models.Payment) *PaymentBuilder {
	builder := new(PaymentBuilder)
	builder.payment = payment
	return builder
}

func (builder *PaymentBuilder) SetId(id int) *PaymentBuilder {
	builder.payment.Id = id
	return builder
}

func (builder *PaymentBuilder) SetPaymentDate() *PaymentBuilder {
	builder.payment.PaymentDate = Helpers.GetCurrentDate()
	return builder
}

func (builder *PaymentBuilder) SetPaymentMode(paymentMode string) *PaymentBuilder {
	builder.payment.PaymentMode = paymentMode
	return builder
}

func (builder *PaymentBuilder) SetAmount(amount float64) *PaymentBuilder {
	builder.payment.Amount = amount
	builder.payment.BaseAmount = amount
	return builder
}

func (builder *PaymentBuilder) SetBalance(balance float64) *PaymentBuilder {
	builder.payment.Balance = balance
	return builder
}

func (builder *PaymentBuilder) SetUsedAmount(usedAmount float64) *PaymentBuilder {
	builder.payment.UsedAmount = usedAmount

	return builder
}

func (builder *PaymentBuilder) SetStatus(status string) *PaymentBuilder {
	builder.payment.Status = status
	return builder
}

func (builder *PaymentBuilder) SetIdCustomer(idCustomer int) *PaymentBuilder {
	builder.payment.IdCustomer = idCustomer
	return builder
}

func (builder *PaymentBuilder) SetDefaults() *PaymentBuilder {
	builder.payment.Type = "customer_payment"
	builder.payment.IdCurrency = 550
	builder.payment.IdChartOfAccounts = 39
	builder.payment.Tag = "3"
	return builder
}
