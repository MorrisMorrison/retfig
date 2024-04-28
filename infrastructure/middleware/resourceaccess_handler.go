package middleware

import (
	"net/http"
	"strings"

	"github.com/MorrisMorrison/retfig/services"
	"github.com/gin-gonic/gin"
)

func ResourceAccessHandler(resourceAccessService *services.ResourceAcessService) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.GetString("currentUser")
		if currentUser == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		eventId := c.Param("eventId")
		presentId := c.Param("presentId")

		canAcessEvents := checkEventAccess(c, resourceAccessService, eventId, currentUser)
		if !canAcessEvents {
			return
		}

		canAccessPresent := checkPresentAccess(c, resourceAccessService, presentId, currentUser)
		if !canAccessPresent {
			return
		}

		c.Next()
	}
}

func checkPresentAccess(c *gin.Context, resourceAccessService *services.ResourceAcessService, presentId string, currentUser string) bool {
	if !pathSegmentExists(c, "presents") {
		return true
	}

	if presentId == "" {
		return true
	}

	allowed, err := resourceAccessService.CanAccessPresent(presentId, currentUser)
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

func checkEventAccess(c *gin.Context, resourceAccessService *services.ResourceAcessService, eventId string, currentUser string) bool {
	if !pathSegmentExists(c, "events") {
		return true
	}

	if eventId == "" {
		return true
	}

	allowed, err := resourceAccessService.CanAccessEvent(eventId, currentUser)
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

func pathSegmentExists(c *gin.Context, pathSegment string) bool {
	pathSegments := strings.Split(c.Request.URL.Path, "/")
	for _, segment := range pathSegments {
		if segment == pathSegment {
			return true
		}
	}

	return false
}
