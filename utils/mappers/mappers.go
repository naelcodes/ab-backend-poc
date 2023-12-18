package mappers

import (
	"encoding/json"
	"strconv"

	"neema.co.za/rest/utils/dto"
	"neema.co.za/rest/utils/models"
)

func CustomerModelToDTO(customer *models.Customer) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = customer.Id
	customerDTO.Customer_name = customer.Customer_name
	customerDTO.State = customer.State
	customerDTO.Account_number = customer.Account_number
	customerDTO.Alias = customer.Alias
	customerDTO.Ab_key = customer.Ab_key
	customerDTO.Tmc_client_number = customer.Tmc_client_number

	return customerDTO
}

func CustomerModelListToDTOList(customers []*models.Customer) []*dto.GetCustomerDTO {
	customerDTOList := make([]*dto.GetCustomerDTO, 0)

	for _, customer := range customers {
		customerDTOList = append(customerDTOList, CustomerModelToDTO(customer))
	}

	return customerDTOList
}

func TravelItemModelToDTO(travelItem *models.TravelItem) *dto.TravelItemDTO {
	travelItemDTO := new(dto.TravelItemDTO)

	travelItemDTO.Id = travelItem.Id
	travelItemDTO.TotalPrice = (*float64)(&travelItem.Total_price)
	travelItemDTO.Itinerary = (*string)(&travelItem.Itinerary)
	travelItemDTO.TravelerName = (*string)(&travelItem.Traveler_name)
	ticketNumber := strconv.Itoa(travelItem.Ticket_number)
	travelItemDTO.TicketNumber = &ticketNumber

	return travelItemDTO
}

func TravelItemModelListToDTOList(travelItems []*models.TravelItem) []*dto.TravelItemDTO {
	travelItemDTOList := make([]*dto.TravelItemDTO, 0)

	for _, travelItem := range travelItems {
		travelItemDTOList = append(travelItemDTOList, TravelItemModelToDTO(travelItem))
	}

	return travelItemDTOList
}

func InvoiceModelToDTO(invoice *models.Invoice, embedCustomer bool) *dto.GetInvoiceDTO {
	invoiceDTO := new(dto.GetInvoiceDTO)
	invoiceDTO.Id = invoice.Id
	invoiceDTO.InvoiceNumber = invoice.Invoice_number
	invoiceDTO.DueDate = invoice.Due_date
	invoiceDTO.CreationDate = invoice.Creation_date
	invoiceDTO.Amount = invoice.Amount
	invoiceDTO.Credit_apply = invoice.Credit_apply
	invoiceDTO.Balance = invoice.Balance
	invoiceDTO.Status = invoice.Status

	if embedCustomer {
		invoiceDTO.IdCustomer = nil
		customer := new(models.Customer)
		_ = json.Unmarshal([]byte(*invoice.Customer), &customer)
		invoiceDTO.Customer = CustomerModelToDTO(customer)
	} else {
		if invoice.Id_customer != nil {
			invoiceDTO.IdCustomer = invoice.Id_customer
		}
	}

	for _, travelItem := range invoice.Travel_items {
		invoiceDTO.TravelItems = append(invoiceDTO.TravelItems, TravelItemModelToDTO(&travelItem))
	}
	return invoiceDTO
}

func InvoiceModelListToDTOList(invoices []*models.Invoice, embedCustomer bool) []*dto.GetInvoiceDTO {
	invoiceDTOList := make([]*dto.GetInvoiceDTO, 0)

	for _, invoice := range invoices {
		invoiceDTOList = append(invoiceDTOList, InvoiceModelToDTO(invoice, embedCustomer))
	}

	return invoiceDTOList
}

func PaymentModelToDTO(payment *models.Payment, embedCustomer bool) *dto.GetPaymentDTO {

	paymentDTO := new(dto.GetPaymentDTO)
	paymentDTO.Id = payment.Id
	paymentDTO.Amount = payment.Amount
	paymentDTO.PaymentNumber = payment.Payment_number
	paymentDTO.PaymentDate = payment.Payment_date
	paymentDTO.PaymentMode = payment.Fop
	paymentDTO.Balance = payment.Balance
	paymentDTO.UsedAmount = payment.Used_amount
	paymentDTO.Status = payment.Status

	if embedCustomer {
		paymentDTO.IdCUstomer = nil
		customer := new(models.Customer)
		_ = json.Unmarshal([]byte(*payment.Customer), &customer)
		paymentDTO.Customer = CustomerModelToDTO(customer)
	} else {
		if payment.Id_customer != nil {
			paymentDTO.IdCUstomer = payment.Id_customer
		}

	}

	return paymentDTO
}

func PaymentModelListToDTOList(payments []*models.Payment, embedCustomer bool) []*dto.GetPaymentDTO {
	paymentDTOList := make([]*dto.GetPaymentDTO, 0)

	for _, payment := range payments {
		paymentDTOList = append(paymentDTOList, PaymentModelToDTO(payment, embedCustomer))
	}

	return paymentDTOList
}
