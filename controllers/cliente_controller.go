package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/simple_api_golang/database"
	"github.com/marcoagpegoraro/simple_api_golang/models"
)

func GetClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente
	database.Instance.Find(&clientes)

	return c.JSON(clientes)
}

func GetCliente(c *fiber.Ctx) error {
	id := c.Params("id")

	var cliente models.Cliente
	database.Instance.First(&cliente, "id = ?", id)

	return c.JSON(cliente)
}

func PostCliente(c *fiber.Ctx) error {
	cliente := new(models.Cliente)

	if err := c.BodyParser(cliente); err != nil {
		return err
	}

	database.Instance.Create(&cliente)

	return c.JSON(cliente)
}
