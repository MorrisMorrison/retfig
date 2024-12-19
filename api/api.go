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

// TODO validate path params (e.g. eventId unused most of the times)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.Index(events.CreateEvent()))
}

func SetTokenCookie(c *gin.Context, username string) {
	tokenExpiresInDuration, err := time.ParseDuration(config.CONFIG.JWTConfig.ExpiresInDuration)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := jwt.GenerateToken(username, config.CONFIG.JWTConfig.Issuer, tokenExpiresInDuration)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		COOKIE_NAME,
		token,
		int(tokenExpiresInDuration/time.Second),
		COOKIE_PATH,
		config.CONFIG.Host,
		true,
		true,
	)
}
