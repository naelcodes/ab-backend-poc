package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/utils/dto"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/types"
)

func (api *Api) GetAllPaymentsHandler(c *fiber.Ctx) error {

	queryParams := c.Locals("queryParams").(*types.GetQueryParams)

	paymentsDTO, err := api.Service.GetAllPaymentsService(queryParams)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting all payments DTO: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("All payments DTO: %v", paymentsDTO))

	return c.Status(fiber.StatusOK).JSON(paymentsDTO)
}

func (api *Api) GetPaymentHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		logger.Error(fmt.Sprintf("Error parsing id: %v", err))
		return CustomErrors.ServiceError(err, "parsing id")
	}

	logger.Info(fmt.Sprintf("params Id: %v", id))

	queryParams := c.Locals("queryParams").(*types.GetQueryParams)

	paymentDTO, err := api.Service.GetPaymentService(id, queryParams)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting payment DTO: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("Payment DTO: %v", paymentDTO))

	return c.Status(fiber.StatusOK).JSON(paymentDTO)
}

func (api *Api) CreatePaymentHandler(c *fiber.Ctx) error {

	createPaymentDTO := c.Locals("payload").(*dto.CreatePaymentDTO)
	logger.Info(fmt.Sprintf("CreatePaymentDTO: %v", createPaymentDTO))

	newPaymentDTO, err := api.Service.CreatePaymentService(createPaymentDTO)

	if err != nil {
		logger.Error(fmt.Sprintf("Error creating payment DTO: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("NewPaymentDTO: %v", newPaymentDTO))
	return c.Status(fiber.StatusCreated).JSON(newPaymentDTO)
}
