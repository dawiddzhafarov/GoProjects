package main

import (
	"github.com/gofiber/fiber"
	"github.com/dawiddzhafarov/GoProjects/GolangFiberCRM/database"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	initDatabase()
	defer database.DBconn
}

func initDatabase() {
	database.
}

func setupRoutes(app *fiber.App) {
	app.Get(GetLead)
	app.Get(GetLeads)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
