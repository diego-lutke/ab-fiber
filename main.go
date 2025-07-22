package main

import (
	"github.com/diego-lutke/ab-fiber/app/middleware"
	"github.com/diego-lutke/ab-fiber/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
