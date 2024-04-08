package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nakamu-ms/hanabi/app/backend/routes"
)

func main() {
	app := fiber.New()
	routes.SetupUserRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
