package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/controllers"
)

func SetupItemRoutes(app *fiber.App) {
	app.Get("/api/items", controllers.GetItems)
	app.Post("/api/items", controllers.CreateItem)
}
