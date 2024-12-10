package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/models"
	"github.com/wbluan/orders-api/src/repository"
)

func GetCustomers(c *fiber.Ctx) error {
	customers, err := repository.GetCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customers)
}

func CreateCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Inputs"})
	}

	id, err := repository.CreateCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	customer.ID = id
	return c.Status(fiber.StatusCreated).JSON(customer)
}
