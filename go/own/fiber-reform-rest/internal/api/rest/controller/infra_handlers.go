package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetPing(ctx *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
	}

	response := Response{Message: "Pong"}
	return ctx.JSON(response)
}
