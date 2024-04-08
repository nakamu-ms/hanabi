package routes

import (
	"github.com/nakamu-ms/hanabi/app/admin/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/api/v1")
	SetupArticles(admin)
}

func SetupArticles(router fiber.Router) {
	articles := router.Group("/articles")
	articles.Get("/", controller.GetArticles)
	articles.Get("/:id", controller.GetArticle)
	articles.Post("/", controller.CreateArticle)
	articles.Put("/:id", controller.UpdateArticle)
	articles.Delete("/:id", controller.DeleteArticle)
}
