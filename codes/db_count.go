package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("can't load or missing .env file")
	}
	app := fiber.New()
	app.Use(logger.New())
	db_url := os.Getenv("DB_URL")
	DB, err = sql.Open("postgres", db_url)
	if err != nil {
		log.Println("can't get driver")
	}
	if err = DB.Ping(); err != nil {
		log.Println("can't connect to the db")
	}
	fmt.Println("db connected sucessfully")
	app.Get("/data", func(c fiber.Ctx) error {
		var count int
		err = DB.QueryRow(`select count(*) from users`).Scan(&count)
		if err != nil {
			return fiber.NewError(500, "can't load")
		}

		return c.JSON(fiber.Map{
			"count": count,
		})

	})

	fmt.Println("server is running")
	if err := app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true}); err != nil {
		log.Fatal(err)
	}
}
