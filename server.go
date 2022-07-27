package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samsul-rijal/go-api/config/database"
	"github.com/samsul-rijal/go-api/config/migration"
	"github.com/samsul-rijal/go-api/routes"
)

func main() {

	// initial database
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	routes.RouteInit(app)

	app.Listen("localhost:5000")
}