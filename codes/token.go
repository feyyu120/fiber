package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

func main() {
	app := fiber.New()

	// Endpoint: /token -> generates and returns a JWT token
	app.Get("/token", func(c fiber.Ctx) error {
		// Define token claims
		claims := jwt.MapClaims{
			"username": "john_doe",
			"role":     "user",
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
		}

		// Create JWT token with HS256 algorithm
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token with secret key
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to generate token",
			})
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"token":  tokenString,
		})
	})

	log.Println("Server is running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
