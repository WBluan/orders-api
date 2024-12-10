package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/controllers"
)

func SetupOrderRoutes(app *fiber.App) {
	app.Get("/api/orders", controllers.GetOrders)
	app.Post("/api/orders", controllers.CreateOrder)
}
