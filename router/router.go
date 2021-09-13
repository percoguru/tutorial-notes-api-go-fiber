package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/percoguru/notes-api-fiber/docs"
	noteRoutes "github.com/percoguru/notes-api-fiber/internals/routes/note"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.Handler)
	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	noteRoutes.SetupNoteRoutes(api)
}
