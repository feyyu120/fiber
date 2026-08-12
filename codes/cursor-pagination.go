package main

import (
	"database/sql"
	"fmt"
	"log"

	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/lib/pq"
	//"github.com/gofiber/fiber/v3/middleware/paginate"
)

type Student struct {
	ID   int
	Name string
}

var DB *sql.DB

func main() {
	var err error
	app := fiber.New()
	app.Use(logger.New())
	db_url := "postgresql://neondb_owner:npg_RTXP8gfVB7pu@ep-small-rain-axw9gkiv-pooler.c-4.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	DB, err = sql.Open("postgres", db_url)

	if err != nil {
		log.Println("error occured")
	}
	if err = DB.Ping(); err != nil {
		log.Println("can't connect to the database")
	}

	log.Println("db connected sunccesfully")

	// start routing
	app.Get("/feed", func(c fiber.Ctx) error {

		after, err := strconv.Atoi(c.Query("after"))
		if err != nil {
			return fiber.ErrBadRequest
		}
		limit := 5
		nextCursor := after + limit + 1
		rows, err := DB.Query(`select id,name from users where id > $1 order by id limit $2`, after, limit)
		if err != nil {
			return fiber.NewError(500, "can't read data")
		}
		var users []Student
		defer rows.Close()
		for rows.Next() {
			var u Student
			if err = rows.Scan(&u.ID, &u.Name); err != nil {
				return fiber.NewError(500, "error when scanning")
			}
			users = append(users, u)
		}
		return c.JSON(fiber.Map{
			"after":      after,
			"data":       users,
			"nextCursor": nextCursor,
		})

	})
	fmt.Println("server is running")
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})
}
