package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var DB *sql.DB

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Order struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
}

func main() {

	var err error
	DB, err = sql.Open("pgx", "postgresql://neondb_owner:npg_RTXP8gfVB7pu@ep-small-rain-axw9gkiv-pooler.c-4.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require")
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Successfully connected to database!")

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/users", func(c fiber.Ctx) error {
		email := c.Query("email")

		// ⏱️ Start a high-precision timer right before the query
		dbStart := time.Now()

		var user User
		err := DB.QueryRow(`SELECT id, name, email FROM users WHERE email=$1`, email).Scan(&user.ID, &user.Name, &user.Email)

		// ⏱️ Stop the timer right after the database responds
		dbDuration := time.Since(dbStart)

		// Print the true isolated database round-trip time to your Go terminal
		log.Printf("[DB TIME] Query for %s took: %v", email, dbDuration)

		if err != nil {
			log.Println("Failed to fetch users:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch users",
			})
		}
		return c.JSON(user)
	})
	app.Get("/ordered", func(c fiber.Ctx) error {
		query, err := DB.Query(`SELECT email,status from users as u JOIN orders as o
ON u.id=o.user_id;`)
		if err != nil {
			log.Println("Failed to fetch users:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch users",
			})
		}
		defer query.Close()

		var results []struct {
			Email  string
			Status string
		}

		for query.Next() {
			var user User
			var order Order
			if err := query.Scan(&user.Email, &order.Status); err != nil {
				return fiber.NewError(500, "something happend")
			}
			results = append(results, struct {
				Email  string
				Status string
			}{
				Email:  user.Email,
				Status: order.Status,
			})
		}
		return c.JSON(results)
	})
	app.Listen(":3000", fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
