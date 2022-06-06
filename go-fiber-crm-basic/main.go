package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/gofiber/fiber"
	"github.com/maitnngo2002/go_lang_projects/go-fiber-crm-basic/database"
	"github.com/maitnngo2002/go_lang_projects/go-fiber-crm-basic/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)

}

func initDatabases() {
	var err error

	// use gorm to connect to your sqlite3 database
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to sqlite3 database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabases()

	setUpRoutes(app)

	app.Listen(3000)
	
	// close the database at the end
	defer database.DBConn.Close() // this line will run after the function has run all other codes
}