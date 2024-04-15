package main

import (
	"fmt"

	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/database/migrations"
	"github.com/MorrisMorrison/retfig/repositories"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/templrender"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Startup retfig")
	dbConn := database.NewConnection()
	migrations.InitializeDatabase(dbConn)

	presentRepository := repositories.NewPresentRepository(dbConn)
	presentService := services.NewPresentService(presentRepository)
	presentAPI := api.NewPresentAPI(presentService)

	eventRepository := repositories.NewEventRepository(dbConn)
	eventService := services.NewEventService(eventRepository, presentService)

	eventAPI := api.NewEventAPI(eventService)

	r := gin.Default()
	r.HTMLRender = &templrender.TemplRender{}

	// r.Use(middleware.HtmxHandler())

	// r.GET("/login", api.GetLogin)
	// r.POST("/login", api.CreateLogin)

	//r.GET("/events/:id", eventAPI.CreateEvent)
	r.POST("/events", eventAPI.CreateEvent)
	r.GET("/events/:id", eventAPI.GetEvent)
	r.DELETE("/events/:id", eventAPI.DeleteEvent)
	r.PATCH("/events/:id", eventAPI.UpdateEvent)

	r.POST("/events/:id/participants", eventAPI.CreateParticipant)
	r.GET("/events/:id/invitation", eventAPI.GetInvitationView)

	r.GET("/events/:id/presents", presentAPI.GetPresents)
	r.POST("/events/:id/presents", presentAPI.CreatePresent)

	r.GET("/", api.Index)

	r.Static("/public", "./public")
	//r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
