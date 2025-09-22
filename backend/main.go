package main

import (
	"log"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"proyecto-universitario/backend/database"
	"proyecto-universitario/backend/routes"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar a la base de datos
	database.ConnectDB()
	
	app := fiber.New()
	
	// Configurar CORS para desarrollo
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})
	
	// Configurar rutas
	routes.SetupRoutes(app)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	log.Fatal(app.Listen(":" + port))
}