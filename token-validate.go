package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("thisismysecret")

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userId,
		"role": "admin",
		"exp":  time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fiber.NewError(500, "can't generate")
	}

	return tokenString, nil

}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/token", func(c fiber.Ctx) error {
		token, err := GenerateToken(3)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})

	})
	log.Println("")
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})
}
