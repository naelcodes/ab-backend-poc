package repository

import (
	"fmt"

	"neema.co.za/rest/modules/payment/internal/domain"
	"neema.co.za/rest/utils/dto"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/mappers"
	dbModels "neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"
)

const tag = "3"
const embedCustomerSqlQuery = "(SELECT to_jsonb(customer) FROM (SELECT id,customer_name,account_number,alias,ab_key,state,tmc_client_number FROM customer WHERE id=payment_received.id_customer) AS customer) AS customer "

const paymentSql = "SELECT id,number,to_char(date,'yyyy-mm-dd') as date,balance::numeric,amount::numeric,used_amount::numeric,fop,status "

func (r *Repository) Count() (*int64, error) {
	logger.Info("Counting payments")

	totalRowCount, err := r.Where("tag = ?", tag).Count(new(dbModels.Payment))

	if err != nil {
		logger.Error(fmt.Sprintf("Error counting payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting payment records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total payment count: %v", totalRowCount))

	return &totalRowCount, nil
}

func (r *Repository) GetAll(queryParams *types.GetQueryParams) (*dto.GetAllPaymentsDTO, error) {

	embedCustomer := false
	paymentSqlQuery := paymentSql

	totalRowCount, err := r.Count()

	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := int(*totalRowCount)

	if queryParams != nil {
		if queryParams.PageNumber != nil && queryParams.PageSize != nil {
			pageNumber = *queryParams.PageNumber
			pageSize = *queryParams.PageSize
		}

		if queryParams.Embed != nil && *queryParams.Embed == "customer" {
			embedCustomer = true
			paymentSqlQuery = paymentSqlQuery + "," + embedCustomerSqlQuery
		} else {
			paymentSqlQuery = paymentSqlQuery + ",id_customer"
		}

	}

	paymentSqlQuery = paymentSqlQuery + " FROM payment_received  WHERE tag = ?  ORDER BY number DESC LIMIT ? OFFSET ?"

	var payments []*dbModels.Payment

	err = r.SQL(paymentSqlQuery, tag, pageSize, pageNumber*pageSize).Find(&payments)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payment records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total payment count: %v", len(payments)))

	paymentDTOList := PaymentModelListToDTOList(payments, embedCustomer)
	return &dto.GetAllPaymentsDTO{
		Data:          paymentDTOList,
		TotalRowCount: int(*totalRowCount),
		PageSize:      pageSize,
		PageNumber:    pageNumber,
	}, nil

}

func (r *Repository) GetById(id int, queryParams *types.GetQueryParams) (*dto.GetPaymentDTO, error) {

	embedCustomer := false
	paymentSqlQuery := paymentSql

	if queryParams != nil {
		if queryParams.Embed != nil && *queryParams.Embed == "customer" {
			embedCustomer = true
			paymentSqlQuery = paymentSqlQuery + "," + embedCustomerSqlQuery
		} else {
			paymentSqlQuery = paymentSqlQuery + ",id_customer"
		}
	}

	paymentSqlQuery = paymentSqlQuery + " FROM payment_received WHERE tag = ? AND id = ?"

	var paymentRecords []*dbModels.Payment

	err := r.SQL(paymentSqlQuery, tag, id).Find(&paymentRecords)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting payment: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payment records: %v", err))
	}

	if len(paymentRecords) == 0 {
		logger.Error(fmt.Sprintf("Payment not found: %v", id))
		return nil, CustomErrors.NotFoundError(fmt.Errorf("payment record not found"))
	}

	logger.Info(fmt.Sprintf("payment count: %v", len(paymentRecords)))

	paymentDTOList := PaymentModelToDTO(paymentRecords[0], embedCustomer)

	return paymentDTOList, nil

}

func (r *Repository) Save(PaymentDomainModel *domain.Payment) (*dto.GetPaymentDTO, error) {

	totalRowCount, err := r.Count()
	if err != nil {
		return nil, err
	}

	payment := PaymentModelFromDomainModel(PaymentDomainModel, int(*totalRowCount))

	_, err = r.Insert(payment)

	if err != nil {
		logger.Error(fmt.Sprintf("Error saving payment: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving payment: %v", err))
	}

	paymentDTO := PaymentModelToDTO(payment, false)

	return paymentDTO, nil
}
