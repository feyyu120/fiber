package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Models struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
