package main

import (
	"os"

	"github.com/Roninors/Expense_Tracker/backend/database"
	"github.com/Roninors/Expense_Tracker/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	utilities.LoadEnv()
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBname:   os.Getenv("DB_NAME"),
	}
	database.ConnectDb(config)
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":8080")
}
