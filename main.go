package main

import (
	"card-master/database"
	config "card-master/lib/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(config.App());
	database.ConnectDB()
	SetupRoutes(app)
	app.Listen(":3000");
}