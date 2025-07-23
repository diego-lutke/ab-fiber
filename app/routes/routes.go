package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/user-agent", func(c *fiber.Ctx) error {
		ua := c.Get("User-Agent")
		return c.SendString(ua)
	})

	app.Get("/status-test", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		// query parameters
		// 0.0.0.0:3000/hello?name=Peter&age=45
		name := c.Query("name")
		age := c.Query("age")

		msg := fmt.Sprintf("%s is %s years old", name, age)
		return c.SendString(msg)
	})

	app.Get("/say/:name/:age/", func(c *fiber.Ctx) error {
		// path parameters
		name := c.Params("name")
		age := c.Params("age")

		msg := fmt.Sprintf("%s is %s years old", name, age)
		return c.SendString(msg)
	})

	app.Get("/now", func(c *fiber.Ctx) error {
		now := time.Now()

		return c.Render("index", fiber.Map{
			"now": now.Format("Jan 2, 2006"),
		})
	})
}
