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

func FindUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User

	database.DB.First(&user, userId)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}
