package routes

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/infrastructure/context"
	"github.com/MorrisMorrison/retfig/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, ctx *context.ApplicationContext) {
	secureEventsAPI := r.Group("/api/htmx/v1/events", middleware.AuthHandler(), middleware.ResourceAccessHandler(ctx.Services.ResourceAcessService), middleware.ViewContextHandler())
	{
		secureEventsAPI.GET("/:eventId", func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.EventAPI.GetEvent)
		})

		secureEventsAPI.GET("/:eventId/invitation", ctx.APIs.ParticipantAPI.GetInvitationView)

		secureEventsAPI.GET("/:eventId/presents", func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.PresentAPI.GetPresents)
		})
		secureEventsAPI.POST("/:eventId/presents", func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.PresentAPI.CreatePresent)
		})

		secureEventsAPI.POST("/:eventId/presents/:presentId/votes", ctx.APIs.VoteAPI.CreateVote)
		secureEventsAPI.POST("/:eventId/presents/:presentId/comments", ctx.APIs.CommentAPI.CreateComment)
		secureEventsAPI.GET("/:eventId/presents/:presentId/comments", ctx.APIs.CommentAPI.GetComments)

		secureEventsAPI.POST("/:eventId/presents/:presentId/claims", func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.ClaimAPI.CreateClaim)
		})
		secureEventsAPI.DELETE("/:eventId/presents/:presentId/claims", func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.ClaimAPI.DeleteClaim)
		})
	}

	publicEventsAPI := r.Group("/events")
	{
		// allows user to reload clear url
		publicEventsAPI.GET("/:eventId", middleware.AuthHandler(), middleware.ResourceAccessHandler(ctx.Services.ResourceAcessService), middleware.ViewContextHandler(), func(c *gin.Context) {
			HandleWithViewContext(c, ctx.APIs.EventAPI.GetEvent)
		})

		publicEventsAPI.GET("/:eventId/invitations", ctx.APIs.ParticipantAPI.GetInvitationView)
		publicEventsAPI.POST("/", ctx.APIs.EventAPI.CreateEvent)
		publicEventsAPI.POST("/:eventId/participants", ctx.APIs.ParticipantAPI.CreateParticipant)
	}

	r.Static("/public", "./ui/public")
	r.GET("/", api.Index)
}
