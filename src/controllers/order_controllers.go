package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/models"
	"github.com/wbluan/orders-api/src/repository"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := repository.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})
	}

	id, err := repository.CreateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	order.ID = id
	return c.Status(fiber.StatusCreated).JSON(order)
}
