package app

import (
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, categoryController controller.CategoryController, customerController controller.CustomerController) {
	authMiddleware := middleware.NewAuthMiddleware()

	api := app.Group("/api", authMiddleware)

	// Category Routes
	categories := api.Group("/categories")
	categories.Post("/", categoryController.Create)
	categories.Put("/:categoryId", categoryController.Update)
	categories.Delete("/:categoryId", categoryController.Delete)
	categories.Get("/:categoryId", categoryController.FindById)
	categories.Get("/", categoryController.FindAll)

	// Customer Routes
	customers := api.Group("/customers")
	customers.Post("/", customerController.Create)
	customers.Put("/:customerId", customerController.Update)
	customers.Delete("/:customerId", customerController.Delete)
	customers.Get("/:customerId", customerController.FindById)
	customers.Get("/", customerController.FindAll)
}