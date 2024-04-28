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

func canAccessEvent(currentUser, eventId string) (bool, error) {
	// true if owner of event of participant of event
	if currentUser == "user123" && eventId == "resource456" {
		return true, nil
	}
	return false, nil
}

func canAccessPresent(currentUser, presentId string) (bool, error) {
	// true if owner of present, owner of event or participant or event
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
