package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/percoguru/notes-api-fiber/database"
	"github.com/percoguru/notes-api-fiber/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
