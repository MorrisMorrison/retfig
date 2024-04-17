package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/views"
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

	voteAPI.voteService.CreateVote(createVoteRequest)

	presentListViewModel, err := voteAPI.presentService.GetPresentListViewModel(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", views.PresentList(*presentListViewModel))
}
