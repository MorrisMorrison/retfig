package routes

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/infrastructure/container"
	"github.com/MorrisMorrison/retfig/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, apis *container.APIContainer, services *container.ServiceContainer) {
	secureEventsAPI := r.Group("/api/htmx/v1/events", middleware.AuthHandler(), middleware.ViewContextHandler())
	{
		secureEventsAPI.GET("/:eventId", func(c *gin.Context) {
			HandleWithViewContext(c, apis.EventAPI.GetEvent)
		})

		secureEventsAPI.GET("/:eventId/invitation", apis.ParticipantAPI.GetInvitationView)

		secureEventsAPI.GET("/:eventId/presents", func(c *gin.Context) {
			HandleWithViewContext(c, apis.PresentAPI.GetPresents)
		})
		secureEventsAPI.POST("/:eventId/presents", func(c *gin.Context) {
			HandleWithViewContext(c, apis.PresentAPI.CreatePresent)
		})

		secureEventsAPI.POST("/:eventId/presents/:presentId/votes", apis.VoteAPI.CreateVote)
		secureEventsAPI.POST("/:eventId/presents/:presentId/comments", apis.CommentAPI.CreateComment)
		secureEventsAPI.GET("/:eventId/presents/:presentId/comments", apis.CommentAPI.GetComments)

		secureEventsAPI.POST("/:eventId/presents/:presentId/claims", func(c *gin.Context) {
			HandleWithViewContext(c, apis.ClaimAPI.CreateClaim)
		})
		secureEventsAPI.DELETE("/:eventId/presents/:presentId/claims", func(c *gin.Context) {
			HandleWithViewContext(c, apis.ClaimAPI.DeleteClaim)
		})
	}

	publicEventsAPI := r.Group("/events")
	{
		// allows user to reload clear url
		publicEventsAPI.GET("/:eventId", middleware.AuthHandler(), middleware.ResourceAccessHandler(services.ResourceAcessService), middleware.ViewContextHandler(), func(c *gin.Context) {
			HandleWithViewContext(c, apis.EventAPI.GetEvent)
		})

		publicEventsAPI.GET("/:eventId/invitations", apis.ParticipantAPI.GetInvitationView)
		publicEventsAPI.POST("/", apis.EventAPI.CreateEvent)
		publicEventsAPI.POST("/:eventId/participants", apis.ParticipantAPI.CreateParticipant)
	}

	r.Static("/public", "./ui/public")
	r.GET("/", api.Index)
}
