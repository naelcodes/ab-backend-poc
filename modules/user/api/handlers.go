package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Api) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching user"})
	}
	return c.JSON(user)
}
