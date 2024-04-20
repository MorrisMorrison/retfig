package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/views/presents"
	"github.com/gin-gonic/gin"
)

type PresentAPI struct {
	presentService services.PresentService
}

func NewPresentAPI(presentService *services.PresentService) *PresentAPI {
	return &PresentAPI{presentService: *presentService}
}

func (presentAPI *PresentAPI) GetPresents(c *gin.Context) {
	eventId := c.Param("eventId")
	presentListViewModel, err := presentAPI.presentService.GetPresentListViewModel(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentList(*presentListViewModel))
}

func (presentAPI *PresentAPI) CreatePresent(c *gin.Context) {
	eventId := c.Param("eventId")
	var createPresentRequest request.CreatePresentRequest

	if err := c.ShouldBindJSON(&createPresentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentAPI.presentService.CreatePresent(eventId, createPresentRequest)

	presentListViewModel, err := presentAPI.presentService.GetPresentListViewModel(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentList(*presentListViewModel))
}
