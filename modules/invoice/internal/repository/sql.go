package repository

import (
	"fmt"

	"neema.co.za/rest/utils/dto"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/mappers"
	dbModels "neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/types"
)

const tag = "3"
const embedCustomerSqlQuery = "(SELECT to_jsonb(customer) FROM (SELECT id,customer_name,account_number,alias,ab_key,state,tmc_client_number FROM customer WHERE id=invoice.id_customer) AS customer) AS customer "

const invoiceSql = "SELECT id,invoice_number,to_char(creation_date,'yyyy-mm-dd') as creation_date,to_char(due_date,'yyyy-mm-dd') as due_date,amount::numeric,balance::numeric,credit_apply::numeric,status,(SELECT jsonb_agg(travel_items) FROM(SELECT id,total_price::numeric,itinerary,traveler_name,ticket_number FROM air_booking WHERE id_invoice=invoice.id) AS travel_items) AS travel_items "

func (r *Repository) Count() (*int64, error) {
	logger.Info("Counting invoices")

	totalRowCount, err := r.Where("tag = ?", tag).Count(new(dbModels.Invoice))

	if err != nil {
		logger.Error(fmt.Sprintf("Error counting invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting invoice records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total invoice count: %v", totalRowCount))

	return &totalRowCount, nil
}

func (r *Repository) GetAll(queryParams *types.GetQueryParams) (*dto.GetAllInvoiceDTO, error) {

	embedCustomer := false
	invoiceQuery := invoiceSql
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
			invoiceQuery = invoiceQuery + "," + embedCustomerSqlQuery
		} else {
			invoiceQuery = invoiceQuery + ",id_customer"
		}

	}

	invoiceQuery = invoiceQuery + " FROM invoice WHERE tag = ? ORDER BY invoice_number ASC LIMIT ? OFFSET ? "

	var invoices []*dbModels.Invoice

	err = r.SQL(invoiceQuery, tag, pageSize, pageNumber*pageSize).Find(&invoices)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoice records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total invoice count: %v", len(invoices)))

	invoiceDTOList := InvoiceModelListToDTOList(invoices, embedCustomer)

	return &dto.GetAllInvoiceDTO{
		Data:          invoiceDTOList,
		TotalRowCount: int(*totalRowCount),
		PageSize:      pageSize,
		PageNumber:    pageNumber,
	}, nil

}

func (r *Repository) GetById(id int, queryParams *types.GetQueryParams) (*dto.GetInvoiceDTO, error) {

	embedCustomer := false
	invoiceQuery := invoiceSql

	if queryParams != nil {
		if queryParams.Embed != nil && *queryParams.Embed == "customer" {
			embedCustomer = true
			invoiceQuery = invoiceQuery + "," + embedCustomerSqlQuery
		} else {
			invoiceQuery = invoiceQuery + ",id_customer"
		}

	}

	invoiceQuery = invoiceQuery + " FROM invoice WHERE id = ?"

	var invoiceRecords []*dbModels.Invoice

	err := r.SQL(invoiceQuery, id).Find(&invoiceRecords)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting invoice: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoice records: %v", err))
	}

	if len(invoiceRecords) == 0 {
		logger.Error(fmt.Sprintf("Error getting invoice: %v", err))
		return nil, CustomErrors.NotFoundError(fmt.Errorf("invoice record not found"))
	}

	logger.Info(fmt.Sprintf("invoice count: %v", len(invoiceRecords)))

	invoiceDTO := InvoiceModelToDTO(invoiceRecords[0], embedCustomer)

	return invoiceDTO, nil

}

func (r *Repository) GetByCustomerId(idCustomer int, queryParams *types.GetQueryParams, paid bool) (*dto.GetCustomerInvoicesDTO, error) {

	WhereCondition := "WHERE tag = ? AND id_customer = ?"
	invoiceQuery := invoiceSql

	if paid {
		WhereCondition = WhereCondition + " AND status = 'paid'"
	} else {
		WhereCondition = WhereCondition + " AND status = 'paid'"
	}

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
	}

	invoiceQuery = invoiceQuery + " FROM invoice " + WhereCondition + " ORDER BY invoice_number ASC LIMIT ? OFFSET ? "

	var customerInvoices []*dbModels.Invoice

	err = r.SQL(invoiceQuery, tag, idCustomer, pageSize, pageNumber*pageSize).Find(&customerInvoices)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer's invoice records: %v", err))
	}

	logger.Info(fmt.Sprintf("Total customer's invoice count: %v", len(customerInvoices)))

	invoiceDTOList := InvoiceModelListToDTOList(customerInvoices, false)
	return &dto.GetCustomerInvoicesDTO{
		Data:          invoiceDTOList,
		TotalRowCount: int(*totalRowCount),
		PageSize:      pageSize,
		PageNumber:    pageNumber,
	}, nil

}
