package api

import (
	"context"

	"neema.co.za/rest/modules/booking/internal/service"
	"neema.co.za/rest/utils/logger"
)

type Exports struct {
	internalService *service.Service
}

func (e *Exports) BM__AssociateInvoiceToTravelItem(context context.Context) (any, error) {
	logger.Info("Associating invoice to travel item")
	return "exported function", nil
}
