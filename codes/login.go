package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Models struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

var validate = validator.New()

func Users() *Models {
	return &Models{Email: "feyyu@gmail", Password: "feysel"}

}

func Login(c fiber.Ctx) error {
	var req Models
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ivalid request",
		})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed the validation",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "logged in succesfully",
		"user": req.Email,
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/login", Login)
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})
}
