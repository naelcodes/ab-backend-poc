package repository

import (
	"fmt"

	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/models"
)

func (r *Repository) GetByInvoiceId(idInvoice int) (any, error) {

	invoices := []*models.Invoice{}
	err := r.SQL("SELECT id,amount::numeric,id_customer FROM invoice WHERE id = ?", idInvoice).Find(&invoices)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoice with id(%v): %v", idInvoice, err))
	}

	if len(invoices) == 0 {
		return nil, CustomErrors.NotFoundError(fmt.Errorf("invoice with id(%v) not found", idInvoice))
	}

	imputationAmountWithRelatedPaymentsQuery := `SELECT i.amount_apply::numeric,p.id as id,p.number ,to_char(p.date,'yyyy-mm-dd') as date,p.balance::numeric,p.amount::numeric AS amount FROM invoice_payment_received AS i RIGHT OUTER JOIN payment_received AS p ON i.id_payment_received=p.id AND id_invoice= ?  WHERE p.id_customer= ? ORDER BY p.number`

	imputationAmountWithRelatedPaymentRecords := make([]*struct {
		Payment       models.Payment `xorm:"extends" json:"payment"`
		AmountApplied float64        `xorm:"amount_apply" json:"amountApplied"`
	}, 0)

	err = r.SQL(imputationAmountWithRelatedPaymentsQuery, idInvoice, invoices[0].IdCustomer).Find(&imputationAmountWithRelatedPaymentRecords)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting imputations of invoice with id(%v) : %v", idInvoice, err))
	}

	if len(imputationAmountWithRelatedPaymentRecords) == 0 {
		return nil, CustomErrors.NotFoundError(fmt.Errorf("imputations of invoice with id(%v) not found", idInvoice))
	}

	data := new(struct {
		InvoiceAmount float64 `json:"invoiceAmount"`
		Imputations   []*struct {
			Payment       models.Payment `xorm:"extends" json:"payment"`
			AmountApplied float64        `xorm:"amount_apply" json:"amountApplied"`
		} `json:"imputations"`
	})

	data.InvoiceAmount = invoices[0].Amount
	data.Imputations = imputationAmountWithRelatedPaymentRecords

	return data, nil
}
