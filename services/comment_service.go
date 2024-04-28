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

func (service *CommentService) GetCommentCountByPresentIds(presentIds []string) (map[string]int32, error) {
	return service.commentRepository.GetCommentCountByPresentIds(presentIds)
}
func (service *CommentService) GetCommentCountByPresentId(presentId string) (int32, error) {
	return service.commentRepository.GetCommentCountByPresentId(uuid.FromStringOrNil(presentId))
}

func (service *CommentService) CreateComment(presentId string, user string, request request.CreateCommentRequest) error {
	comment := service.mapCreateCommentRequestToComment(presentId, user, request)
	return service.commentRepository.CreateComment(comment)
}

func (service *CommentService) mapCreateCommentRequestToComment(presentId string, user string, request request.CreateCommentRequest) models.Comment {
	return *models.NewComment(uuid.FromStringOrNil(presentId), request.Content, user)
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

	return viewmodels.NewCommentListViewModel(eventId, presentId, commentListItems), nil
}

func (service *CommentService) GetCommentListItemViewModel(presentId string, username string, content string) (*viewmodels.CommentListItemViewModel, error) {
	return viewmodels.NewCommentListItemViewModel(presentId, username, content, ""), nil
}

func (service *CommentService) mapCommentToCommentListItemViewModel(comment *models.Comment) *viewmodels.CommentListItemViewModel {
	dateLayout := "January 02, 2006"
	return viewmodels.NewCommentListItemViewModel(comment.PresentId.String(), comment.CreatedBy, comment.Content, comment.CreatedAt.Format(dateLayout))
}
