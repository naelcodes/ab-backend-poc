package api

import (
	"github.com/gofiber/fiber/v2"
)

func (this *Api) GetUserByID(c *fiber.Ctx) error {
	//id, _ := strconv.Atoi(c.Params("id"))
	user, err := this.Service.GetUserByID(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching user"})
	}
	return c.JSON(user)
}
