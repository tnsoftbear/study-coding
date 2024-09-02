package router

import (
	"fiber-reform-rest/internal/api/rest/controller"
	"fiber-reform-rest/internal/infra/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gopkg.in/reform.v1"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App, reformDB *reform.DB) {
	newsRepo := storage.NewNewsRepository(reformDB)
	api := app.Group("", limiter.New())
	app.Use(logger.New())
	api.Get("/ping", controller.GetPing)
	api.Get("/list", func(ctx *fiber.Ctx) error { return controller.GetNewsList(ctx, newsRepo) })
	api.Post("/edit/:Id", func(ctx *fiber.Ctx) error { return controller.PostNewsEditById(ctx, newsRepo) })
	api.Post("/add", func(ctx *fiber.Ctx) error { return controller.PostNewsAdd(ctx, newsRepo) })
	api.Post("/add-category/:NewsId/:CatId", func(ctx *fiber.Ctx) error { return controller.PostNewsAddCategory(ctx, newsRepo) })
}
