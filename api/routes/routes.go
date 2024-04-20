package routes

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/infrastructure/container"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, apis *container.APIContainer) {
	r.POST("/events", apis.EventAPI.CreateEvent)
	r.GET("/events/:eventId", apis.EventAPI.GetEvent)
	r.DELETE("/events/:eventId", apis.EventAPI.DeleteEvent)
	r.PATCH("/events/:eventId", apis.EventAPI.UpdateEvent)

	r.POST("/events/:eventId/participants", apis.ParticipantAPI.CreateParticipant)
	r.GET("/events/:eventId/invitation", apis.ParticipantAPI.GetInvitationView)

	r.GET("/events/:eventId/presents", apis.PresentAPI.GetPresents)
	r.POST("/events/:eventId/presents", apis.PresentAPI.CreatePresent)
	r.POST("/events/:eventId/presents/:presentId/vote", apis.VoteAPI.CreateVote)
	r.POST("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.CreateComment)
	r.GET("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.GetComments)

	r.Static("/public", "./ui/public")
	r.GET("/", api.Index)
}
