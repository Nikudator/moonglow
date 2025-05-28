package routes

import (
	"../handlers"
	"github.com/gofiber/fiber"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/posts", handlers.GetProducts)
	api.Post("/posts", handlers.CreateProduct)
	api.Get("/posts/:id", handlers.GetProduct)
	api.Put("/posts/:id", handlers.UpdateProduct)
	api.Delete("/posts/:id", handlers.DeleteProduct)
}
