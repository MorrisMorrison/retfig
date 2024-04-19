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

	voteRepository := repositories.NewVoteRepository(dbConn)
	voteService := services.NewVoteService(voteRepository)
	commentRepository := repositories.NewCommentRepository(dbConn)
	commentService := services.NewCommentService(commentRepository)
	presentRepository := repositories.NewPresentRepository(dbConn)
	presentService := services.NewPresentService(presentRepository, voteService, commentService)
	presentAPI := api.NewPresentAPI(presentService)
	commentAPI := api.NewCommentAPI(commentService, presentService)
	eventRepository := repositories.NewEventRepository(dbConn)
	eventService := services.NewEventService(eventRepository, presentService)
	eventAPI := api.NewEventAPI(eventService)
	voteAPI := api.NewVoteAPI(voteService, presentService)

	r := gin.Default()
	r.HTMLRender = &templrender.TemplRender{}

	// r.Use(middleware.HtmxHandler())

	// r.GET("/login", api.GetLogin)
	// r.POST("/login", api.CreateLogin)

	//r.GET("/events/:id", eventAPI.CreateEvent)
	r.POST("/events", eventAPI.CreateEvent)
	r.GET("/events/:eventId", eventAPI.GetEvent)
	r.DELETE("/events/:eventId", eventAPI.DeleteEvent)
	r.PATCH("/events/:eventId", eventAPI.UpdateEvent)

	r.POST("/events/:eventId/participants", eventAPI.CreateParticipant)
	r.GET("/events/:eventId/invitation", eventAPI.GetInvitationView)

	r.GET("/events/:eventId/presents", presentAPI.GetPresents)
	r.POST("/events/:eventId/presents", presentAPI.CreatePresent)

	r.POST("/events/:eventId/presents/:presentId/vote", voteAPI.CreateVote)
	
	r.POST("/events/:eventId/presents/:presentId/comments", commentAPI.CreateComment)
	r.GET("/events/:eventId/presents/:presentId/comments", commentAPI.GetComments)

	r.GET("/", api.Index)

	r.Static("/public", "./public")
	//r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
