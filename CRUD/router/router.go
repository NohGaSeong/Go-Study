package router

import (
	"golang-fiber-crud/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome go",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/note/:noteId", func(router fiber.Router) {
		router.Delete("", noteController.Delete)
		router.Get("", noteController.FIndById)
		router.Patch("", noteController.Update)
	})

	return router
}
