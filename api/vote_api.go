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
	eventId := c.Param("eventId")

	var createVoteRequest request.CreateVoteRequest

	if err := c.ShouldBindJSON(&createVoteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := voteAPI.voteService.CreateVote(createVoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	voteButtonsViewModel, err := voteAPI.voteService.GetVoteButtonsViewModel(eventId, createVoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "", votes.VoteButtons(eventId, voteButtonsViewModel))
}
