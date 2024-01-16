package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/views"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.Index())
}
