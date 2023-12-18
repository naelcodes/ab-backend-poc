package domain

import (
	"errors"
	"fmt"

	CustomErrors "neema.co.za/rest/utils/errors"
	Helpers "neema.co.za/rest/utils/helpers"
	"neema.co.za/rest/utils/types"
)

type Payment struct {
	types.BaseEntity
	PaymentNumber string
	PaymentDate   string
	PaymentMode   string
	Amount        float64
	Balance       float64
	UsedAmount    float64
	Status        string
	IdCustomer    int
}

func (p *Payment) calculateBalance() error {

	if p.UsedAmount > p.Amount {
		return CustomErrors.DomainError(errors.New("payment balance can't be less than 0"))
	}

	p.Balance = Helpers.RoundDecimalPlaces(p.Amount-p.UsedAmount, 2)
	// p.updateStatus()

	return nil
}

func (p *Payment) updateStatus() {

	if p.UsedAmount == p.Amount && p.Balance == 0 {
		p.Status = "used"
	} else {
		p.Status = "open"
	}
}

func (p *Payment) AllocateAmount(imputationAmount float64) error {

	// if p.Status == "used" {
	// 	return CustomErrors.DomainError(fmt.Errorf("payment %v is already used. new allocations can't be made on a used payment", p.PaymentNumber))
	// }

	if p.UsedAmount+imputationAmount > p.Amount {
		return CustomErrors.DomainError(fmt.Errorf("allocated(used) amount on payment %v can't be greater than the payment amount", p.PaymentNumber))
	}

	p.UsedAmount = p.UsedAmount + Helpers.RoundDecimalPlaces(imputationAmount, 2)
	err := p.calculateBalance()

	if err != nil {
		return err
	}

	return nil
}
