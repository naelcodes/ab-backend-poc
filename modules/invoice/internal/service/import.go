package service

import (
	"context"

	"neema.co.za/rest/utils/managers"
)

type Imports struct {
	dependencyManager *managers.DependencyManager
}

func (i *Imports) AssociateInvoiceToTravelItems(IdInvoice int, travelItemIds []int) (any, error) {
	fn := i.dependencyManager.Get("AssociateInvoiceToTravelItems")
	return fn(context.Background())
}
