package service

import (
	"context"

	"neema.co.za/rest/utils/managers"
	"neema.co.za/rest/utils/types"
	"xorm.io/xorm"
)

type Imports struct {
	dependencyManager *managers.DependencyManager
}

func (i *Imports) AssociateInvoiceToTravelItems(requestContext context.Context, transaction *xorm.Session, IdInvoice int, travelItemIds []int) (any, error) {
	AssociateInvoiceToTravelItems := i.dependencyManager.Get("BM__AssociateInvoiceToTravelItems")

	requestContext = context.WithValue(requestContext, types.KeyType("invoiceId"), IdInvoice)
	requestContext = context.WithValue(requestContext, types.KeyType("travelItemIds"), travelItemIds)
	requestContext = context.WithValue(requestContext, types.KeyType("transaction"), transaction)

	return AssociateInvoiceToTravelItems(requestContext)
}
