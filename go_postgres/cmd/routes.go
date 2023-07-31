package main

import (
	handlers "command-line-argumentsC:\\Users\\Adam\\Desktop\\gooooo\\go_exercise\\go_postgres\\handlers\\facts.go"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home(c))
}
