package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nakamu-ms/hanabi/app/admin/routes"
)

func main() {
	app := fiber.New()
	routes.SetupAdminRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
