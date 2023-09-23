package main

import (
	"fmt"
	_ "invoice-app/database"
	"invoice-app/routes"
	"invoice-app/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
}

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)

	appPort := utils.GetPort()

	err := app.Listen(fmt.Sprintf(":%s", appPort))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
