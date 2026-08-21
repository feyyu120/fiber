package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Users struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"requred,email"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/create", func(c fiber.Ctx) error {
		var req Users
		err := c.Bind().Body(&req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "created sucessfully",
		})
	})

}
