package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/healtcheck", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	log.Fatal(app.Listen(":4000"))
}
