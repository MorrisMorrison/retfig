package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/viewcontext"
	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/MorrisMorrison/retfig/ui/views/events"
	"github.com/MorrisMorrison/retfig/utils/links"

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

	eventId, err := eventAPI.eventService.CreateEvent(createEventRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	viewModel, err := eventAPI.eventService.GetEventViewModel(eventId.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	// used to include id of created event in url so users can hit reload
	c.Header("HX-Push-Url", links.BuildGetEventLink(eventId.String()))

	SetTokenCookie(c, createEventRequest.Username)

	vc := viewcontext.NewViewContext(createEventRequest.Username, true)
	c.HTML(http.StatusOK, "", events.GetEvent(vc, viewModel))
}

func (eventAPI *EventAPI) GetEvent(c *gin.Context, vc *viewcontext.ViewContext) {
	eventId := c.Param("eventId")
	viewModel, err := eventAPI.eventService.GetEventViewModel(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "", views.Index(events.GetEvent(vc, viewModel)))
}
