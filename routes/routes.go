package routes

import (
	"testlogin/controllers"
	"github.com/gofiber/fiber/v2")

func SetRoutes(app *fiber.App){
	app.Get("/", controllers.GetUsers)
	app.Get("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}