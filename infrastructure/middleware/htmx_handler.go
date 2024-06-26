package middleware

import (
	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/MorrisMorrison/retfig/ui/views/events"

	"github.com/gin-gonic/gin"
)

func HtmxHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		isHxRequest := c.GetHeader("HX-Request")

		if isHxRequest == "" || isHxRequest == "false" {
			c.HTML(204, "", views.Index(events.CreateEvent()))
			c.Abort()
			return
		}

		c.Next()
	}
}
