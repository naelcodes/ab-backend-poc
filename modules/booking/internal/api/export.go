package api

import (
	"context"

	"neema.co.za/rest/modules/booking/internal/service"
	"xorm.io/xorm"
)

type Exports struct {
	internalService *service.Service
}

func (e *Exports) BM__AssociateInvoiceToTravelItems(context context.Context) (any, error) {

	transaction := context.Value("transaction").(*xorm.Session)
	invoiceId := context.Value("invoiceId").(int)
	travelItemIds := context.Value("travelItemIds").([]int)
	err := e.internalService.AddInvoiceToTravelItems(transaction, invoiceId, travelItemIds)
	return nil, err
}
