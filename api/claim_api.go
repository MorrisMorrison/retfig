package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/viewcontext"
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

func (claimAPI *ClaimAPI) CreateClaim(c *gin.Context, vc *viewcontext.ViewContext) {
	currentUser := c.GetString("currentUser")
	presentId := c.Param("presentId")

	err := claimAPI.claimService.CreateClaim(presentId, currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentListItemViewModel, err := claimAPI.presentService.GetPresentListItemViewModel(presentId, currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentListItem(vc, presentListItemViewModel))
}

func (claimAPI *ClaimAPI) DeleteClaim(c *gin.Context, vc *viewcontext.ViewContext) {
	currentUser := c.GetString("currentUser")
	presentId := c.Param("presentId")

	err := claimAPI.claimService.DeleteClaim(presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presentListItemViewModel, err := claimAPI.presentService.GetPresentListItemViewModel(presentId, currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", presents.PresentListItem(vc, presentListItemViewModel))
}
