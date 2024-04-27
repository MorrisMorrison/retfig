package middleware

import (
	"github.com/MorrisMorrison/retfig/ui/viewcontext"
	"github.com/gin-gonic/gin"
)

func ViewContextHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.GetString("currentUser")
		viewContext := viewcontext.NewViewContext(currentUser, currentUser != "")

		c.Set("viewContext", viewContext)

		c.Next()
	}
}
