package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResourceAccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.GetString("currentUser")
		if currentUser == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		eventId := c.Param("eventId")
		presentId := c.Param("presentId")

		canAcessEvents := checkEventAccess(c, eventId, currentUser)
		if !canAcessEvents {
			return
		}

		canAccessPresent := checkPresentAccess(c, presentId, currentUser)
		if !canAccessPresent {
			return
		}

		c.Next()
	}
}

func checkPresentAccess(c *gin.Context, presentId string, currentUser string) bool {
	if !pathSegmentExists(c, "presents") {
		return true
	}

	allowed, err := canAccessPresent(currentUser, presentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify resource access"})
		return false
	}

	if !allowed {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return false
	}

	return true
}

func checkEventAccess(c *gin.Context, eventId string, currentUser string) bool {
	if !pathSegmentExists(c, "events") {
		return true
	}

	allowed, err := canAccessEvent(currentUser, eventId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify resource access"})
		return false
	}

	if !allowed {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return false
	}

	return true
}

func canAccessEvent(currentUser, eventId string) (bool, error) {
	if currentUser == "user123" && eventId == "resource456" {
		return true, nil
	}
	return false, nil
}

func canAccessPresent(currentUser, presentId string) (bool, error) {
	if currentUser == "user123" && presentId == "resource456" {
		return true, nil
	}
	return false, nil
}

func pathSegmentExists(c *gin.Context, pathSegment string) bool {
	pathSegments := strings.Split(c.Request.URL.Path, "/")
	for _, segment := range pathSegments {
		if segment == pathSegment {
			return true
		}
	}

	return false
}
