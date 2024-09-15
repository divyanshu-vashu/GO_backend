package main

import (
	"21BRS1444_backend/database"
	"21BRS1444_backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Replace with your frontend URL
		AllowCredentials: true, // Enable credentials (cookies, HTTP authentication, etc.)
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))
	app.Static("/uploads", "./uploads")
	
	routes.Setup(app)

	app.Listen(":8000")
}
