package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/views/votes"
	"github.com/gin-gonic/gin"
)

type VoteAPI struct {
	voteService    services.VoteService
	presentService services.PresentService
}

func NewVoteAPI(voteService *services.VoteService, presentService *services.PresentService) *VoteAPI {
	return &VoteAPI{voteService: *voteService, presentService: *presentService}
}

func (voteAPI *VoteAPI) CreateVote(c *gin.Context) {
	currentUser := c.GetString(PARAM_CURRENT_USER)
	eventId := c.Param(PARAM_EVENT_ID)
	presentId := c.Param(PARAM_PRESENT_ID)

	var createVoteRequest request.CreateVoteRequest

	if err := c.ShouldBindJSON(&createVoteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	err := voteAPI.voteService.CreateVote(presentId, currentUser, createVoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	voteButtonsViewModel, err := voteAPI.voteService.GetVoteButtonsViewModel(eventId, presentId, currentUser, createVoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "", votes.VoteButtons(voteButtonsViewModel))
}
