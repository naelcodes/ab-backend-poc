package service

import (
	"fmt"

	"neema.co.za/rest/modules/payment/internal/domain"
	"neema.co.za/rest/utils/dto"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/types"
)

func (s *Service) GetAllPaymentsService(queryParams *types.GetQueryParams) (*dto.GetAllPaymentsDTO, error) {
	logger.Info("Getting all payments")
	return s.Repository.GetAll(queryParams)
}

func (s *Service) GetPaymentService(id int, queryParams *types.GetQueryParams) (*dto.GetPaymentDTO, error) {
	logger.Info("Getting payment")
	return s.Repository.GetById(id, queryParams)
}

func (s *Service) CreatePaymentService(paymentDTO *dto.CreatePaymentDTO) (*dto.GetPaymentDTO, error) {
	logger.Info("Creating payment")

	paymentBuilder := domain.NewPaymentBuilder().
		SetAmount(paymentDTO.Amount).
		SetIdCustomer(paymentDTO.IdCustomer).
		SetPaymentMode(paymentDTO.PaymentMode).
		SetPaymentDate().
		SetBalance(paymentDTO.Amount).
		SetStatus("open").
		SetUsedAmount(0)

	err := paymentBuilder.Validate()

	if err != nil {
		logger.Error(fmt.Sprintf("payment validation error: %v", err))
		return nil, err
	}

	paymentDomainModel := paymentBuilder.Build()

	return s.Repository.Save(paymentDomainModel)
}
