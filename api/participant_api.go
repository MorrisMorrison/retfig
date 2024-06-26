package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/utils/links"

	"github.com/MorrisMorrison/retfig/ui/viewcontext"
	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/MorrisMorrison/retfig/ui/views/events"
	"github.com/MorrisMorrison/retfig/ui/views/invitations"

	"github.com/gin-gonic/gin"
)

type ParticipantAPI struct {
	participantService services.ParticipantService
	eventService       services.EventService
}

func NewParticipantAPI(participantService *services.ParticipantService, eventService *services.EventService) *ParticipantAPI {
	return &ParticipantAPI{participantService: *participantService, eventService: *eventService}
}

func (participantAPI *ParticipantAPI) CreateParticipant(c *gin.Context) {
	var createParticipantRequest request.CreateParticipantRequest

	currentUser := c.GetString(PARAM_CURRENT_USER)
	eventId := c.Param(PARAM_EVENT_ID)

	if err := c.ShouldBindJSON(&createParticipantRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	err := participantAPI.participantService.CreateParticipant(eventId, createParticipantRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	viewModel, err := participantAPI.eventService.GetEventViewModel(eventId, currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return

	}

	c.Header(HEADER_HX_PUSH_URL, links.BuildGetEventLink(eventId))

	SetTokenCookie(c, createParticipantRequest.Username)

	vc := viewcontext.NewViewContext(createParticipantRequest.Username, true)

	c.HTML(http.StatusOK, "", events.GetEvent(vc, viewModel))
}

func (eventAPI *ParticipantAPI) GetInvitationView(c *gin.Context) {
	eventId := c.Param(PARAM_EVENT_ID)

	c.HTML(http.StatusOK, "", views.Index(invitations.GetInvitation(eventId)))
}
