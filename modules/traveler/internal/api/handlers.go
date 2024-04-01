package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/utils/logger"
	"neema.co.za/rest/utils/payloads"
	"neema.co.za/rest/utils/types"
)

func (api *Api) CreateTravelerHandler(c *fiber.Ctx) error {
	CreateTravelerPayload := c.Locals("payload").(*payloads.CreateTravelerPayload)

	logger.Info(fmt.Sprintf("CreateTravelerDTO: %v", CreateTravelerPayload))

	newTravelerRecord, err := api.CreateTravelerService(*CreateTravelerPayload)

	if err != nil {
		logger.Error(fmt.Sprintf("Error creating traveler record: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("NewTravelerRecord: %v", newTravelerRecord))
	return c.Status(fiber.StatusCreated).JSON(newTravelerRecord)
}

func (api *Api) GetSingleTravelerHandler(c *fiber.Ctx) error {
	return nil
}

func (api *Api) GetAllTravelersHandler(c *fiber.Ctx) error {
	queryParams := c.Locals("queryParams").(*types.GetQueryParams)
	travelers, err := api.GetAllTravelersService(queryParams)

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting all travelers : %v", err))
		return err
	}

	return c.Status(fiber.StatusOK).JSON(travelers)
}

func (api *Api) UpdateTravelerHandler(c *fiber.Ctx) error {
	return nil
}

func (api *Api) DeleteTravelerHandler(c *fiber.Ctx) error {
	return nil
}
