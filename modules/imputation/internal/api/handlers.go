package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	CustomErrors "neema.co.za/rest/utils/errors"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/payloads"
)

func (api *Api) GetImputationsHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		logger.Error(fmt.Sprintf("Error parsing id: %v", err))
		return CustomErrors.ServiceError(err, "parsing id")
	}

	logger.Info(fmt.Sprintf("params Id: %v", id))

	invoiceImputationRecords, err := api.Service.GetImputationsService(id)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting invoice imputations: %v", err))
		return err
	}

	return c.Status(fiber.StatusOK).JSON(invoiceImputationRecords)

}

func (api *Api) ApplyImputationsHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		logger.Error(fmt.Sprintf("Error parsing id: %v", err))
		return CustomErrors.ServiceError(err, "parsing id")
	}

	logger.Info(fmt.Sprintf("params Id: %v", id))

	payload := c.Locals("payload").([]*payloads.ImputationPayload)

	// logic

	return c.Status(fiber.StatusOK).JSON(payload)

}
