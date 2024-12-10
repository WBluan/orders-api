package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	SetupCustomerRoutes(app)
	SetupItemRoutes(app)
	SetupOrderRoutes(app)
}
