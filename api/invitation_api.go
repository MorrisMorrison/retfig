package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvitationAPI struct {
}

func NewInvitationAPI() *InvitationAPI {
	return &InvitationAPI{}
}

func (invitationAPI *InvitationAPI) GetInvitationView(c *gin.Context) {
	eventId := c.Param("id")
	c.HTML(http.StatusOK, "", eventId)
}
