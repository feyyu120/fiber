package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Users struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
}

var validate = validator.New()

func Create(c fiber.Ctx) error {
	var req Users
	err := c.Bind().Body(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err = validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "created sucessfully",
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/create", Create)
	log.Println("server is running")
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})
}
