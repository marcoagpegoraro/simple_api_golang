package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/simple_api_golang/controllers"
	"github.com/marcoagpegoraro/simple_api_golang/database"
)

func init() {
	database.InitializeDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/clientes", controllers.GetClientes)
	app.Get("/clientes/:id", controllers.GetCliente)
	app.Post("/clientes", controllers.PostCliente)

	app.Listen(":3000")
}
