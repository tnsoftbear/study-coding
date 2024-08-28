package bootstrap

import (
	"fiber-reform-rest/internal/api/rest/router"
	"fiber-reform-rest/internal/infra/env"
	"fiber-reform-rest/internal/infra/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp() *fiber.App {
	env.Setup()
	reformDB := storage.Setup()
	// app := fiber.New(fiber.Config{})
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.Setup(app, reformDB)
	return app
}
