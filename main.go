// main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"neema.co.za/rest/api/handlers"
	"neema.co.za/rest/database"
	"neema.co.za/rest/repository"
	"neema.co.za/rest/service"
)

func main() {
	app := fiber.New()

	// Initialize PostgreSQL database
	engine := database.InitDB()

	// Repository, Service, and Handler initialization
	userRepository := repository.NewUserRepository(engine)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService, app.Group)
	userHandler.Start()
	// Start the Fiber app
	log.Fatal(app.Listen(":8080"))
}
