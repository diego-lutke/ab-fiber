package main

import (
	"github.com/diego-lutke/ab-fiber/app/middleware"
	"github.com/diego-lutke/ab-fiber/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(middleware.Logger())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
