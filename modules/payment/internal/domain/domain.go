package domain

import (
	"errors"
	"fmt"

	CustomErrors "neema.co.za/rest/utils/errors"
	Helpers "neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/models"
)

type PaymentDomain struct {
	payment *models.Payment
	builder *PaymentBuilder
	errors  error
}

func NewPaymentDomain() *PaymentDomain {
	domain := &PaymentDomain{payment: &models.Payment{}}
	return domain

}

func (domain *PaymentDomain) GetPaymentBuilder() *PaymentBuilder {
	if domain.builder == nil {
		domain.builder = NewPaymentBuilder(domain.payment)
	}
	return domain.builder
}

func (domain *PaymentDomain) GetPayment() *models.Payment {
	return domain.payment
}

func (domain *PaymentDomain) calculateBalance() error {

	if domain.payment.UsedAmount > domain.payment.Amount {
		return CustomErrors.DomainError(errors.New("payment balance can't be less than 0"))
	}

	domain.payment.Balance = Helpers.RoundDecimalPlaces(domain.payment.Amount-domain.payment.UsedAmount, 2)
	domain.updateStatus()

	return nil
}

func (domain *PaymentDomain) updateStatus() {

	if domain.payment.UsedAmount == domain.payment.Amount && domain.payment.Balance == 0 {
		domain.payment.Status = "used"
	} else {
		domain.payment.Status = "open"
	}
}

func (domain *PaymentDomain) AllocateAmount(imputationAmount float64) error {

	// if p.Status == "used" {
	// 	return CustomErrors.DomainError(fmt.Errorf("payment %v is already used. new allocations can't be made on a used payment", p.PaymentNumber))
	// }

	if domain.payment.UsedAmount+imputationAmount > domain.payment.Amount {
		return CustomErrors.DomainError(fmt.Errorf("allocated(used) amount on payment %v can't be greater than the payment amount", domain.payment.PaymentNumber))
	}

	domain.payment.UsedAmount = domain.payment.UsedAmount + Helpers.RoundDecimalPlaces(imputationAmount, 2)
	err := domain.calculateBalance()

	if err != nil {
		return err
	}

	return nil
}

func (domain *PaymentDomain) Validate() error {

	if domain.payment.Amount < 0 {
		domain.errors = errors.Join(domain.errors, fmt.Errorf("payment.amount can't be less than 0"))
	}

	if domain.payment.Balance < 0 {
		domain.errors = errors.Join(domain.errors, fmt.Errorf("payment.balance can't be less than 0"))
	}

	if domain.payment.UsedAmount < 0 {
		domain.errors = errors.Join(domain.errors, fmt.Errorf("payment.usedAmount can't be less than 0"))
	}

	if domain.payment.Balance != Helpers.RoundDecimalPlaces(domain.payment.Amount-domain.payment.UsedAmount, 2) {
		domain.errors = errors.Join(domain.errors, fmt.Errorf("payment.balance (%v) must be equal to the difference between payment.amount (%v) and payment.usedAmount (%v)", domain.payment.Balance, float64(domain.payment.Amount), float64(domain.payment.UsedAmount)))
	}

	if domain.payment.UsedAmount > domain.payment.Amount {
		domain.errors = errors.Join(domain.errors, fmt.Errorf("payment.usedAmount can't be greater than payment.amount"))
	}

	if domain.errors != nil {
		return CustomErrors.DomainError(domain.errors)
	}
	return nil

}
