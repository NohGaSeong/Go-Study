package controller

import (
	"golang-fiber-crud/data/request"
	"golang-fiber-crud/data/response"
	"golang-fiber-crud/helper"
	"golang-fiber-crud/service"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	err := ctx.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "success",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
