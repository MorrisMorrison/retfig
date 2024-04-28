package api

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/services"
	"github.com/MorrisMorrison/retfig/ui/views/comments"

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
	eventId := c.Param(PARAM_EVENT_ID)
	presentId := c.Param(PARAM_PRESENT_ID)

	commentListViewModel, err := commentAPI.commentService.GetCommentListViewModel(eventId, presentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", comments.CommentList(commentListViewModel))
}

func (commentAPI *CommentAPI) CreateComment(c *gin.Context) {
	currentUser := c.GetString(PARAM_CURRENT_USER)
	presentId := c.Param(PARAM_PRESENT_ID)

	var createCommentRequest request.CreateCommentRequest
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	err := commentAPI.commentService.CreateComment(presentId, currentUser, createCommentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	commentListItemViewModel, err := commentAPI.commentService.GetCommentListItemViewModel(presentId, currentUser, createCommentRequest.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
	}

	c.HTML(http.StatusOK, "", comments.CommentListItem(commentListItemViewModel))
}
