package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	articles := []string{"Article 1", "Article 2", "Article 3"}

	return c.JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	articles := []string{"Article 1", "Article 2", "Article 3"}

	return c.JSON(articles)
}

func CreateArticle(c *fiber.Ctx) error {
	article := new(struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	})
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Request body is not valid")
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	article := new(struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	})
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Request body is not valid")
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	article := new(struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	})
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Request body is not valid")
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}
