package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hikmahnurarifah/webbackend/database"
	"github.com/hikmahnurarifah/webbackend/routes"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{":" + port}, // Sesuaikan dengan origin frontend Anda
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Authorization",
	}))

	routes.Setup(app)
	app.Listen(":" + port)

}
