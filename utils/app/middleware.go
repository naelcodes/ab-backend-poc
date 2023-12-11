package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func getCors() func(*fiber.Ctx) error {
	config := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}

	return cors.New(config)
}

func getRecover() func(*fiber.Ctx) error {
	config := recover.Config{EnableStackTrace: true}
	return recover.New(config)
}
