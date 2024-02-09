package main

import (
	"apresentacao_go/controllers"
	"apresentacao_go/database"

	"github.com/gofiber/fiber/v2"
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
