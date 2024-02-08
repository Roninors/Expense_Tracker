package router

import (
	"github.com/Roninors/Expense_Tracker/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Post("/register", handler.RegisterUser)
	v1.Get("/:id", handler.FindUser)
}
