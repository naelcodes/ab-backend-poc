package service

import "neema.co.za/rest/utils/payloads"

func (s *Service) GetImputationsService(idInvoice int) (any, error) {
	return s.Repository.GetByInvoiceId(idInvoice)
}

func (s *Service) ApplyImputationsService(idInvoice int, payload []*payloads.ImputationPayload) (int, int, int, error) {
	return 0, 0, 0, nil
}
