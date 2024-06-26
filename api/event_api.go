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
	currentUser := c.GetString(PARAM_CURRENT_USER)

	if err := c.ShouldBindJSON(&createEventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	eventId, err := eventAPI.eventService.CreateEvent(createEventRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	viewModel, err := eventAPI.eventService.GetEventViewModel(eventId.String(), currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return

	}

	// used to include id of created event in url so users can hit reload
	c.Header(HEADER_HX_PUSH_URL, links.BuildGetEventLink(eventId.String()))

	SetTokenCookie(c, createEventRequest.Username)

	vc := viewcontext.NewViewContext(createEventRequest.Username, true)
	c.HTML(http.StatusOK, "", events.GetEvent(vc, viewModel))
}

func (eventAPI *EventAPI) GetEvent(c *gin.Context, vc *viewcontext.ViewContext) {
	currentUser := c.GetString(PARAM_CURRENT_USER)
	eventId := c.Param(PARAM_EVENT_ID)

	viewModel, err := eventAPI.eventService.GetEventViewModel(eventId, currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "", views.Index(events.GetEvent(vc, viewModel)))
}
