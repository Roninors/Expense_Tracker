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

	return c.Status(200).JSON(newUser)
}

func FindUser(c *fiber.Ctx) error {
	email := c.Params("email")
	var user models.User

	database.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintln("user not found")})
	}

	return c.Status(200).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	var user models.Credentials
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "failed to parse body"})
	}
	var userFound models.User
	database.DB.Where("email = ?", user.Email).First(&userFound)
	if userFound.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintln("user not found")})
	}
	if match := utilities.CheckHashPassword(user.Password, userFound.Password); !match {
		return c.Status(401).JSON(fiber.Map{"error": fmt.Sprintln("Invalid Credentials")})
	}
	t, exp, err := utilities.GenerateJWT(userFound)
	if err != nil {
		panic(err)
	}
	return c.Status(200).JSON(fiber.Map{"token": t, "expires": exp})
}
