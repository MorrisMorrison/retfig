package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/views/presents"

	"github.com/gin-gonic/gin"
)

type ClaimAPI struct {
	claimService   services.ClaimService
	presentService services.PresentService
}

func NewClaimAPI(claimService *services.ClaimService, presentService *services.PresentService) *ClaimAPI {
	return &ClaimAPI{claimService: *claimService, presentService: *presentService}
}

func (claimAPI *ClaimAPI) CreateClaim(c *gin.Context) {
	// eventId := c.Param("eventId")
	presentId := c.Param("presentId")

	var createClaimRequest request.CreateClaimRequest

	if err := c.ShouldBindJSON(&createClaimRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := claimAPI.claimService.CreateClaim(presentId, createClaimRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentListItemViewModel, err := claimAPI.presentService.GetPresentListItemViewModel(presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentListItem(presentListItemViewModel))
}

func (claimAPI *ClaimAPI) DeleteClaim(c *gin.Context) {
	// eventId := c.Param("eventId")
	presentId := c.Param("presentId")

	err := claimAPI.claimService.DeleteClaim(presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentListItemViewModel, err := claimAPI.presentService.GetPresentListItemViewModel(presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentListItem(presentListItemViewModel))
}
