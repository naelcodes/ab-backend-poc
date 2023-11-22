// handlers/user_handler.go
package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"neema.co.za/rest/service"
)

type UserHandler struct {
	userService   *service.UserService
	RouterCreator func(prefix string, handlers ...fiber.Handler) fiber.Router
}

func NewUserHandler(userService *service.UserService, RouterCreator func(prefix string, handlers ...fiber.Handler) fiber.Router) *UserHandler {
	return &UserHandler{userService, RouterCreator}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.userService.GetUserByID(id)
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
