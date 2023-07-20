package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type JSONTextResponse struct {
	Message string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello Go Fiber!")

		return c.JSON(JSONTextResponse{Message: "Hello Fiber"})
	})

	log.Fatal(app.Listen(":3000"))
}
