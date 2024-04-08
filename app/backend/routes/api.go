package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/api/v1")
	user.Get("/", getUserInfo)
	user.Put("/", updateUserInfo)
}

func getUserInfo(c *fiber.Ctx) error {
	return c.SendString("User information")
}

func updateUserInfo(c *fiber.Ctx) error {
	return c.SendString("User information updated")
}
