package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/config"
	"github.com/wbluan/orders-api/src/repository"
	"github.com/wbluan/orders-api/src/routes"
)

func main() {
	repository.Init()

	app := fiber.New()
	routes.SetupRoutes(app)
	port := config.GetPort()
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
