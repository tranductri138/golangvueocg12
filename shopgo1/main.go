package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	product "github.com/trisdepzai/go-fiber-tutorial/models"
	user "github.com/trisdepzai/go-fiber-tutorial/models"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to TrisdepzaiHome!!")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)

	app.Get("/products", product.GetProducts)
	app.Get("/product/:id", product.GetProduct)
	app.Post("/product", product.SaveProduct)
	app.Delete("/product/:id", product.DeleteProduct)
	app.Put("/product/:id", product.UpdateProduct)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")
}
