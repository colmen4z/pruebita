package handlers

import (
	"github.com/gofiber/fiber/v2"
	"proyecto-universitario/backend/database"
	"proyecto-universitario/backend/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"message": "Error al cargar usuarios",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Usuarios cargados",
		"data": users,
		"count": len(users),
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"message": "Usuario no encontrado",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Usuario encontrado",
		"data": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	
	database.DB.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"message": "Usuario no encontrado",
		})
	}

	result = database.DB.Delete(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"message": "Error al eliminar al usuario",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Usuario eliminado con exito",
		"data": user,
	})
}