package main

import (
	"log"
	"testlogin/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"testlogin/database"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.SetRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
