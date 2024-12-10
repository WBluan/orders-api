package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/orders-api/src/models"
	"github.com/wbluan/orders-api/src/repository"
)

func GetItems(c *fiber.Ctx) error {
	items, err := repository.GetItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error})
	}
	return c.JSON(items)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	id, err := repository.CreateItem(item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	item.ID = id
	return c.Status(fiber.StatusCreated).JSON(item)
}
