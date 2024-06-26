package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/payloads"
	"neema.co.za/rest/utils/types"
)

func (api *Api) GetAllCustomerHandler(c *fiber.Ctx) error {

	queryParams := c.Locals("queryParams").(*types.GetQueryParams)
	customersDTO, err := api.Service.GetAllCustomerService(queryParams)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting all customers DTO: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("All customers DTO: %v", customersDTO))
	return c.Status(fiber.StatusOK).JSON(customersDTO)
}

func (api *Api) GetCustomerHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		logger.Error(fmt.Sprintf("Error parsing id: %v", err))
		return CustomErrors.ServiceError(err, "parsing id")
	}

	logger.Info(fmt.Sprintf("params Id: %v", id))
	customerDTO, err := api.Service.GetCustomerService(id)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting customer DTO: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("Customer DTO: %v", customerDTO))
	return c.Status(fiber.StatusOK).JSON(customerDTO)
}

func (api *Api) CreateCustomerHandler(c *fiber.Ctx) error {

	CreateCustomerPayload := c.Locals("payload").(*payloads.CreateCustomerPayload)
	logger.Info(fmt.Sprintf("CreateCustomerDTO: %v", CreateCustomerPayload))

	newCustomerDTO, err := api.Service.CreateCustomerService(*CreateCustomerPayload)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating customer DTO: %v", err))
		return err
	}
	logger.Info(fmt.Sprintf("NewCustomerDTO: %v", newCustomerDTO))
	return c.Status(fiber.StatusCreated).JSON(newCustomerDTO)

}

func (api *Api) UpdateCustomerHandler(c *fiber.Ctx) error {

	customerID, _ := c.ParamsInt("id")
	updateCustomerPayload := c.Locals("payload").(*payloads.UpdateCustomerPayload)
	logger.Info(fmt.Sprintf("UpdateCustomerDTO: %v", updateCustomerPayload))

	err := api.Service.UpdateCustomerService(customerID, *updateCustomerPayload)
	if err != nil {
		logger.Error(fmt.Sprintf("Error updating customer DTO: %v", err))
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("Customer with id %v updated", customerID))
}

func (api *Api) DeleteCustomerHandler(c *fiber.Ctx) error {
	customerID, _ := c.ParamsInt("id")
	err := api.Service.DeleteCustomerService(customerID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting customer DTO: %v", err))
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("Customer with id %v deleted", customerID))
}
