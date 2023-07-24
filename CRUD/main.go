package main

import (
	"fmt"
	"golang-fiber-crud/config"
	"golang-fiber-crud/controller"
	"golang-fiber-crud/model"
	"golang-fiber-crud/repository"
	"golang-fiber-crud/router"
	"golang-fiber-crud/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("Run service ... ")

	loadConfig, err := config.Loadconfig(",")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	//database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	//Init Repo
	noteRepository := repository.NewNoteRepositoryImpl(db)

	//Init service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	//Init Controller
	noteController := controller.NewNoteController(noteService)

	//Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)
	log.Fatal(app.Listen(":8000"))
}
