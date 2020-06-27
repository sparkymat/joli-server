package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) {
		c.Send(fmt.Sprintf("Hello %s!", c.Params("name")))
	})

	app.Listen(3000)
}
