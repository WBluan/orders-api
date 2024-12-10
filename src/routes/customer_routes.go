package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/controllers"
)

func SetupCustomerRoutes(app *fiber.App) {
	app.Get("/api/customers", controllers.GetCustomers)
	app.Post("/api/customers", controllers.CreateCustomer)
}