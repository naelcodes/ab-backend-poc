package service

import (
	"context"

	"xorm.io/xorm"
)

type Exports struct {
	InternalService *Service
}

func (e *Exports) BM__AssociateInvoiceToTravelItems(context context.Context) (any, error) {

	transaction := context.Value("transaction").(*xorm.Session)
	invoiceId := context.Value("invoiceId").(int)
	travelItemIds := context.Value("travelItemIds").([]int)
	err := e.InternalService.AddInvoiceToTravelItems(transaction, invoiceId, travelItemIds)
	return nil, err
}
