package main

import (
	"log"

	"github.com/Inigojeevan/fiber-GORM/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to fiber")
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Get("/api", welcome)
	log.Fatal(app.Listen(":3000"))
}
