package main

import (
	"github.com/MorrisMorrison/retfig/context"
	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/persistence/database/migrations"
	"github.com/MorrisMorrison/retfig/routes"
	"github.com/MorrisMorrison/retfig/templrender"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := database.NewConnection()
	migrations.InitializeDatabase(dbConn)

	applicationContext := context.NewApplicationContext()

	r := gin.Default()
	r.HTMLRender = &templrender.TemplRender{}

	routes.ConfigureRoutes(r, applicationContext.APIs)

	r.Run()
}
