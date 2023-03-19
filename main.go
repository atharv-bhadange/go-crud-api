package main

import (
	"github.com/atharv-bhadange/go-crud-api/user"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Router(app *fiber.App) {
	app.Get("/users", user.GetUsersList)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Get("/user/:id", user.GetUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Router(app)
	app.Listen(":3000")
}
