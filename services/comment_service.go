package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/persistence/models"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	"github.com/MorrisMorrison/retfig/ui/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type CommentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(commentRepository *repositories.CommentRepository) *CommentService {
	return &CommentService{commentRepository: *commentRepository}
}

func (service *CommentService) GetCommentCountMapByPresentIds(presentIds []string) (map[string]int32, error) {
	return service.commentRepository.GetCommentCountMapByPresentIds(presentIds)
}

func (service *CommentService) CreateComment(request request.CreateCommentRequest) error {
	comment := service.mapCreateCommentRequestToComment(request)
	return service.commentRepository.CreateComment(comment)
}

func (service *CommentService) mapCreateCommentRequestToComment(request request.CreateCommentRequest) models.Comment {
	createdUpdated := models.CreatedUpdated{
		CreatedBy: request.Username,
		UpdatedBy: request.Username,
	}

	return models.Comment{
		PresentId:      uuid.FromStringOrNil(request.PresentId),
		Content:        request.Content,
		CreatedUpdated: createdUpdated,
	}
}

func (service *CommentService) GetCommentListViewModel(eventId string, presentId string) (*viewmodels.CommentListViewModel, error) {
	var commentListItems []*viewmodels.CommentListItemViewModel
	comments, err := service.commentRepository.GetCommentsByPresentId(uuid.FromStringOrNil(presentId))
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		commentListItem := service.mapCommentToCommentListItemViewModel(comment)
		commentListItems = append(commentListItems, commentListItem)
	}

	return &viewmodels.CommentListViewModel{
		PresentId: presentId,
		Comments:  commentListItems,
	}, nil
}

func (service *CommentService) GetCommentListItemViewModel(username string, content string) (*viewmodels.CommentListItemViewModel, error) {
	return &viewmodels.CommentListItemViewModel{
		Username:  username,
		Content:   content,
		CreatedAt: "",
	}, nil
}

func (service *CommentService) mapCommentToCommentListItemViewModel(comment *models.Comment) *viewmodels.CommentListItemViewModel {
	dateLayout := "January 02, 2006"

	return &viewmodels.CommentListItemViewModel{
		Username:  comment.CreatedBy,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(dateLayout),
	}
}
