package controller

import (
	"fmt"
	"strconv"

	"fiber-reform-rest/internal/domain/model"
	"fiber-reform-rest/internal/domain/repository"

	"github.com/gofiber/fiber/v2"
)

func GetNewsList(c *fiber.Ctx, repo repository.NewsRepository) error {
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per-page", 10)
	news := repo.Load(page, perPage)
	return c.JSON(fiber.Map{
		"Success": true,
		"News":    news,
	})
}

func PostNewsEditById(c *fiber.Ctx, repo repository.NewsRepository) error {
	id, err := strconv.Atoi(c.Params("Id"))
	news := &model.News{
		ID: int64(id),
		Title: "xxx",
		Content: "ccc xxx",
	}
	repo.Save(news)
	if err != nil {
		return c.SendStatus(404)
	}
	return c.SendString(fmt.Sprintf("Id is %d", id))
}
