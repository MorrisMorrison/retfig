package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/views"
	"github.com/gin-gonic/gin"
)

type EventAPI struct {
	eventService services.EventService
}

func NewEventAPI(eventService *services.EventService) *EventAPI {
	return &EventAPI{eventService: *eventService}
}

func (eventAPI *EventAPI) CreateEvent(c *gin.Context) {
	var createEventRequest request.CreateEventRequest

	if err := c.ShouldBindJSON(&createEventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createEventResponse, err := eventAPI.eventService.CreateEvent(createEventRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentType := c.GetHeader("Content-Type")
	if contentType == "text/html" {
		c.HTML(http.StatusOK, "", views.GetEvent(createEventResponse))
		return
	}

	c.JSON(http.StatusCreated, createEventResponse)
}
