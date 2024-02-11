package handler

import (
	"fmt"

	"github.com/Roninors/Expense_Tracker/backend/database"
	"github.com/Roninors/Expense_Tracker/backend/models"
	"github.com/Roninors/Expense_Tracker/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var newUser models.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "failed to parse body"})
	}
	hashedPassword, err := utilities.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Failed to hash password: %v", err)})
	}
	newUser.Password = hashedPassword
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to create user: %v", result.Error)})
	}

	return c.JSON(newUser)
}

func FindUser(c *fiber.Ctx) error {
	email := c.Params("email")
	var user models.User

	database.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}
