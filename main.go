package main

import (
	"github.com/MorrisMorrison/retfig/api/routes"
	"github.com/MorrisMorrison/retfig/infrastructure/context"
	"github.com/MorrisMorrison/retfig/infrastructure/middleware"
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
	r.Use(middleware.ViewContextHandler())
	routes.ConfigureRoutes(r, applicationContext.APIs)

	r.Run()
}
