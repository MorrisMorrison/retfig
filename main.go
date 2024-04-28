package main

import (
	"github.com/MorrisMorrison/retfig/infrastructure/context"
	"github.com/MorrisMorrison/retfig/infrastructure/routes"
	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/persistence/database/migrations"
	"github.com/MorrisMorrison/retfig/ui/templrender"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := database.NewConnection()
	migrations.InitializeDatabase(dbConn)

	applicationContext := context.NewApplicationContext()

	r := gin.Default()
	r.HTMLRender = &templrender.TemplRender{}
	routes.ConfigureRoutes(r, applicationContext.APIs, applicationContext.Services)

	r.Run()
}
