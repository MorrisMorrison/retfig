package middleware

import (
	"fmt"

	"github.com/MorrisMorrison/retfig/ui/views"
	"github.com/gin-gonic/gin"
)

func HtmxHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		isHxRequest := c.GetHeader("HX-Request")

		fmt.Println(isHxRequest)
		if isHxRequest == "" || isHxRequest == "false" {
			c.HTML(204, "", views.Index(views.CreateEvent()))
			c.Abort()
			return
		}

		c.Next()
	}
}
