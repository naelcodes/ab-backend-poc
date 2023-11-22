// handlers/user_handler.go
package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"neema.co.za/rest/app"
	"neema.co.za/rest/service"
)

type UserHandler struct {
	*service.UserService
	app.RouterCreator
}

func NewUserHandler(userService *service.UserService, routerCreator app.RouterCreator) *UserHandler {
	return &UserHandler{userService, routerCreator}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching user"})
	}
	return c.JSON(user)
}

func (h *UserHandler) Start() {
	router := h.RouterCreator("/api/v1")
	h.handleCustomerRoutes(router.Group("/customers"))
}

func (h *UserHandler) handleCustomerRoutes(router fiber.Router) {
	router.Get("", h.GetUserByID)
}
