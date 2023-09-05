package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-laravel/app/controllers"
	"github.com/HerlambangHaryo/go-crud-laravel/app/middleware"
)

	// ApiRoutes function to define api routes
		func ApiRoutes(app *fiber.App) {
			// Create a group for api routes with prefix /api
			api := app.Group("/api")

			// Create a group for user routes with prefix /user and auth middleware
			user := api.Group("/user", middleware.AuthMiddleware)

			// Register the user create handler with POST method and path /
			user.Post("/", controllers.UserCreate)

			// Register the user read handler with GET method and path /:id
			user.Get("/:id", controllers.UserRead)

			// Register the user update handler with PUT method and path /:id
			user.Put("/:id", controllers.UserUpdate)

			// Register the user delete handler with DELETE method and path /:id
			user.Delete("/:id", controllers.UserDelete)
		}