package routes

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/ui/viewcontext"
	"github.com/gin-gonic/gin"
)

type ViewContextHandlerFunc func(*gin.Context, *viewcontext.ViewContext)

func HandleWithViewContext(c *gin.Context, handler ViewContextHandlerFunc) {
	vc, exists := c.Get("viewContext")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ViewContext not found"})
		return
	}

	viewContext, ok := vc.(*viewcontext.ViewContext)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ViewContext data"})
		return
	}

	handler(c, viewContext)
}
