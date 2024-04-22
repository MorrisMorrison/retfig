package routes

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/infrastructure/container"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, apis *container.APIContainer) {
	a := r.Group("/api")
	{
		htmx := a.Group("/htmx")
		{
			v1 := htmx.Group("/v1")
			{
				v1.POST("/events", apis.EventAPI.CreateEvent)
				v1.GET("/events/:eventId", apis.EventAPI.GetEvent)
				v1.DELETE("/events/:eventId", apis.EventAPI.DeleteEvent)
				v1.PATCH("/events/:eventId", apis.EventAPI.UpdateEvent)

				v1.POST("/events/:eventId/participants", apis.ParticipantAPI.CreateParticipant)
				v1.GET("/events/:eventId/invitation", apis.ParticipantAPI.GetInvitationView)

				v1.GET("/events/:eventId/presents", apis.PresentAPI.GetPresents)
				v1.POST("/events/:eventId/presents", apis.PresentAPI.CreatePresent)

				v1.POST("/events/:eventId/presents/:presentId/votes", apis.VoteAPI.CreateVote)
				v1.POST("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.CreateComment)
				v1.GET("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.GetComments)

				v1.POST("/events/:eventId/presents/:presentId/claim", apis.PresentAPI.ClaimPresent)
				v1.DELETE("/events/:eventId/presents/:presentId/claim", apis.PresentAPI.UnclaimPresent)

			}
		}
	}

	r.GET("/events/:eventId", apis.EventAPI.GetEvent)
	r.GET("/events/:eventId/invitations", apis.ParticipantAPI.GetInvitationView)

	r.Static("/public", "./ui/public")
	r.GET("/", api.Index)
}
