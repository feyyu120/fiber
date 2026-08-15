package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *sql.DB

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("can't load or missing .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgresql://neondb_owner:npg_RTXP8gfVB7pu@ep-small-rain-axw9gkiv-pooler.c-4.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	}

	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer DB.Close()

	if err = DB.Ping(); err != nil {
		log.Fatal("can't ping database:", err)
	}
	log.Println("database connected successfully")

	app := fiber.New()
	app.Use(logger.New())

	// Offset Pagination Route
	app.Get("/users", func(c fiber.Ctx) error {
		page, err := strconv.Atoi(c.Query("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		var total int
		err = DB.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&total)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to count users"})
		}

		rows, err := DB.Query(`SELECT id, name, email FROM users ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "failed to scan user"})
			}
			users = append(users, u)
		}

		totalPages := int(math.Ceil(float64(total) / float64(limit)))

		return c.JSON(fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
			"data":        users,
		})
	})

	fmt.Println("server is running on :3000")
	if err := app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true}); err != nil {
		log.Fatal(err)
	}
}
