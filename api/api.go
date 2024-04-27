package api

import (
	"net/http"
	"time"

	"github.com/MorrisMorrison/retfig/infrastructure/jwt"
	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/MorrisMorrison/retfig/ui/views/events"
	"github.com/gin-gonic/gin"
)

const TOKEN_EXPIRES_IN = time.Hour * 24

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.Index(events.CreateEvent()))
}

func SetTokenCookie(c *gin.Context, username string) {
	token, err := jwt.GenerateToken(username, TOKEN_EXPIRES_IN)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie(
		"token",
		token,
		int(TOKEN_EXPIRES_IN/time.Second),
		"/",
		"127.0.0.1",
		false,
		true,
	)
}
