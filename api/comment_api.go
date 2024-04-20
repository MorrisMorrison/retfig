package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/views"
	"github.com/gin-gonic/gin"
)

type CommentAPI struct {
	commentService services.CommentService
	presentService services.PresentService
}

func NewCommentAPI(commentService *services.CommentService, presentService *services.PresentService) *CommentAPI {
	return &CommentAPI{commentService: *commentService, presentService: *presentService}
}

func (commentAPI *CommentAPI) GetComments(c *gin.Context) {
	eventId := c.Param("eventId")
	presentId := c.Param("presentId")

	commentListViewModel, err := commentAPI.commentService.GetCommentListViewModel(eventId, presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", views.CommentList(eventId, commentListViewModel))
}

func (commentAPI *CommentAPI) CreateComment(c *gin.Context) {
	eventId := c.Param("eventId")
	presentId := c.Param("presentId")

	var createCommentRequest request.CreateCommentRequest

	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := commentAPI.commentService.CreateComment(createCommentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentListViewModel, err := commentAPI.commentService.GetCommentListViewModel(eventId, presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", views.CommentList(eventId, commentListViewModel))
}