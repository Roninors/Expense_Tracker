package handler

import (
	"github.com/Roninors/Expense_Tracker/backend/database"
	"github.com/Roninors/Expense_Tracker/backend/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var newUser models.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to parse body"})
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(newUser)
}

func Test(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"mssg": "hello world!"})
}
