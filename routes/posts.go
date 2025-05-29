package routes

import (
	"moonglow/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/posts", handlers.GetPosts)
	api.Post("/posts", handlers.CreatePost)
	api.Get("/posts/:id", handlers.GetPost)
	api.Put("/posts/:id", handlers.UpdatePost)
	api.Delete("/posts/:id", handlers.DeletePost)
}
