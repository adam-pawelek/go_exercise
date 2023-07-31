package main

import (
	"github.com/adam-pawelek/go_exercise/blob/main/go_postgres/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home(c))
}
