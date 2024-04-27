package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/viewcontext"
	"github.com/MorrisMorrison/retfig/ui/views/presents"
	"github.com/gin-gonic/gin"
)

type PresentAPI struct {
	presentService services.PresentService
}

func NewPresentAPI(presentService *services.PresentService) *PresentAPI {
	return &PresentAPI{presentService: *presentService}
}

func (presentAPI *PresentAPI) GetPresents(c *gin.Context, vc *viewcontext.ViewContext) {
	eventId := c.Param("eventId")
	presentListViewModel, err := presentAPI.presentService.GetPresentListViewModel(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.HTML(http.StatusOK, "", presents.PresentList(vc, presentListViewModel))
}

func (presentAPI *PresentAPI) CreatePresent(c *gin.Context, vc *viewcontext.ViewContext) {
	var createPresentRequest request.CreatePresentRequest

	currentUser := c.GetString("currentUser")
	eventId := c.Param("eventId")

	if err := c.ShouldBindJSON(&createPresentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentId, err := presentAPI.presentService.CreatePresent(eventId, currentUser, createPresentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentListItemViewModel, err := presentAPI.presentService.GetSimplePresentListItemViewModel(presentId.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentListItem(vc, presentListItemViewModel))
}
