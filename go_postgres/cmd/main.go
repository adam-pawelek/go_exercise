package main

import (
	"github.com/adam-pawelek/go_exercise/tree/main/go_postgres/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
