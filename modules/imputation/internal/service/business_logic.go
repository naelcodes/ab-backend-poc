package service

func (s *Service) GetImputationsService(idInvoice int) (any, error) {
	return s.Repository.GetByInvoiceId(idInvoice)
}
