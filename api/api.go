package api

import (
	"net/http"
	"time"

	"github.com/MorrisMorrison/retfig/infrastructure/config"
	"github.com/MorrisMorrison/retfig/infrastructure/jwt"
	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/MorrisMorrison/retfig/ui/views/events"
	"github.com/gin-gonic/gin"
)

// TODO validata path params

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.Index(events.CreateEvent()))
}

func SetTokenCookie(c *gin.Context, username string) {
	tokenExpiresIn := config.GetEnv(config.CONFIG_KEY_JWT_EXPIRES_IN_DURATION, "24h")
	tokenExpiresInDuration, err := time.ParseDuration(tokenExpiresIn)
	if err != nil {
		c.Error(err)
		return
	}

	issuer := config.GetEnv(config.CONFIG_KEY_JWT_ISSUER, "retfig.com")
	token, err := jwt.GenerateToken(username, issuer, tokenExpiresInDuration)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie(
		"token",
		token,
		int(tokenExpiresInDuration/time.Second),
		"/",
		"127.0.0.1",
		false,
		true,
	)
}
