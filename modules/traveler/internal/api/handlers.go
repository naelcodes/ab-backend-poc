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

	logger.Info(fmt.Sprintf("CreateTravelerPayload: %v", CreateTravelerPayload))

	newTravelerRecord, err := api.CreateTravelerService(*CreateTravelerPayload)

	if err != nil {
		logger.Error(fmt.Sprintf("Error creating traveler record: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("NewTravelerRecord: %v", newTravelerRecord))
	return c.Status(fiber.StatusCreated).JSON(newTravelerRecord)
}

func (api *Api) GetTravelerByIdHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	traveler, err := api.GetTravelerByIdService(id)
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting traveler: %v", err))
		return err
	}
	return c.Status(fiber.StatusOK).JSON(traveler)
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
	travelerId, _ := c.ParamsInt("id")
	UpdateTravelerPayload := c.Locals("payload").(*payloads.UpdateTravelerPayload)

	logger.Info(fmt.Sprintf("UpdateTravelerPayload: %v", UpdateTravelerPayload))

	updatedTraveler, err := api.UpdateTravelerService(travelerId, *UpdateTravelerPayload)

	if err != nil {
		logger.Error(fmt.Sprintf("Error updating traveler: %v", err))
		return err
	}

	return c.Status(fiber.StatusOK).JSON(updatedTraveler)
}

func (api *Api) DeleteTravelerHandler(c *fiber.Ctx) error {
	return nil
}
