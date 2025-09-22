package routes

import (
	"github.com/gofiber/fiber/v2"
	"proyecto-universitario/backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/users")
	{
		users.Get("/", handlers.GetUsers)
		users.Get("/:id", handlers.GetUser)
		users.Post("/", handlers.CreateUser)
		users.Delete("/:id", handlers.DeleteUser)
	}
	
	//api.Get("/users", handlers.GetUsers)
	//api.Post("/users", handlers.CreateUser)
}