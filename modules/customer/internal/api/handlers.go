package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/utils/logger"
)

func (api *Api) GetAllCustomerHandler(c *fiber.Ctx) error {

	user, err := api.Service.GetAllCustomerService()
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting all customers: %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching user"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
