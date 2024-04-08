package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/api/v1")
	admin.Get("/", getAdminInfo)
	admin.Put("/", updateAdminInfo)
}

func getAdminInfo(c *fiber.Ctx) error {
	return c.SendString("Admin information")
}

func updateAdminInfo(c *fiber.Ctx) error {
	// 管理者情報の更新処理をここに書く
	return c.SendString("Admin information updated")
}
