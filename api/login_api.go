package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/views"
	"github.com/gin-gonic/gin"
)

func CreateLogin(c *gin.Context) {
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.Login())
}
