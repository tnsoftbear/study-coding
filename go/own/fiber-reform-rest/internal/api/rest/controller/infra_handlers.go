package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetPing(c *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
	}

	response := Response{Message: "Pong"}
	return c.JSON(response)
}
