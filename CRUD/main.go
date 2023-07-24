package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	app := fiber.New()

	app.Mount("/api", router)

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to golang , fiber, and GOrm",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
