package api

import (
	"fmt"
	"net/http"

	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/views"
	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var createEventRequest models.CreateEventRequest

	if err := c.ShouldBindJSON(&createEventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event := models.Event{
		Name:      createEventRequest.Name,
		Owner:     createEventRequest.Username,
		Recipient: createEventRequest.Recipient,
		Users:     []string{},
	}

	c.HTML(http.StatusOK, "", views.GetEvent(event))
}

func GetEvents(c *gin.Context) {
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}
