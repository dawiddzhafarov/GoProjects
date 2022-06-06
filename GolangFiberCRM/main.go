package main

import (
	"C"
	"fmt"
	"github.com/dawiddzhafarov/GoProjects/GolangFiberCRM/database"
	"github.com/dawiddzhafarov/GoProjects/GolangFiberCRM/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	initDatabase()
	defer database.DBConn.Close()
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to databse")
	}
	fmt.Println("Connection with database succesful")
	database.DBConn.AutoMigrate((&lead.Lead{}))
	fmt.Println("Database migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}
