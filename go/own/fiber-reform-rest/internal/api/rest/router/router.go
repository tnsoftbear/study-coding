package router

import (
	"fiber-reform-rest/internal/api/rest/controller"
	"fiber-reform-rest/internal/infra/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gopkg.in/reform.v1"
)

func Setup(app *fiber.App, reformDB *reform.DB) {
	newsRepo := storage.NewNewsRepository(reformDB)
	api := app.Group("", limiter.New())
	api.Get("/ping", controller.GetPing)
	api.Get("/list", func(ctx *fiber.Ctx) error { return controller.GetNewsList(ctx, newsRepo) })
	api.Post("/edit/:Id", func(ctx *fiber.Ctx) error { return controller.PostNewsEditById(ctx, newsRepo) })
}
