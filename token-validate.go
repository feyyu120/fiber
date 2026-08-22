package main

import (
	"log"
	"strings"
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

func ValidateToken(c fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" || !strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		return fiber.NewError(401, "token missing")
	}

	token := strings.SplitN(auth, " ", 2)
	verified, err := jwt.Parse(token[1], func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil || !verified.Valid {
		return fiber.NewError(403, "invalid token")
	}

	claims, ok := verified.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(401, "invalid token format")
	}

	c.Locals("userId", claims)
	return c.Next()

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
	app.Get("/profile", ValidateToken, func(c fiber.Ctx) error {
		user := c.Locals("userId")

		return c.JSON(fiber.Map{
			"msg":  "user is logged in",
			"user": user,
		})

	})

	log.Println("")
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})
}
