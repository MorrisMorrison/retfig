package routes

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/infrastructure/container"
	"github.com/MorrisMorrison/retfig/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, apis *container.APIContainer) {
	a := r.Group("/api")
	{
		htmx := a.Group("/htmx")
		{
			v1 := htmx.Group("/v1")
			{
				v1.Use(middleware.AuthHandler())

				v1.GET("/events/:eventId", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.EventAPI.GetEvent)
				})

				v1.POST("/events/:eventId/participants", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.ParticipantAPI.CreateParticipant)
				})

				v1.GET("/events/:eventId/invitation", apis.ParticipantAPI.GetInvitationView)

				v1.GET("/events/:eventId/presents", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.PresentAPI.GetPresents)
				})
				v1.POST("/events/:eventId/presents", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.PresentAPI.CreatePresent)
				})

				v1.POST("/events/:eventId/presents/:presentId/votes", apis.VoteAPI.CreateVote)
				v1.POST("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.CreateComment)
				v1.GET("/events/:eventId/presents/:presentId/comments", apis.CommentAPI.GetComments)

				v1.POST("/events/:eventId/presents/:presentId/claims", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.ClaimAPI.CreateClaim)
				})
				v1.DELETE("/events/:eventId/presents/:presentId/claims", func(c *gin.Context) {
					api.HandleWithViewContext(c, apis.ClaimAPI.DeleteClaim)
				})
			}
		}
	}

	r.GET("/events/:eventId/invitations", apis.ParticipantAPI.GetInvitationView)
	r.POST("/events", func(c *gin.Context) {
		api.HandleWithViewContext(c, apis.EventAPI.CreateEvent)
	})

	r.Static("/public", "./ui/public")
	r.GET("/", api.Index)
}
