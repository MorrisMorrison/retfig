package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/infrastructure/logger"
	"github.com/MorrisMorrison/retfig/persistence/models"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	"github.com/MorrisMorrison/retfig/ui/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type PresentService struct {
	presentRepository repositories.PresentRepository
	voteService       VoteService
	commentService    CommentService
}

func NewPresentService(presentRepository *repositories.PresentRepository, voteService *VoteService, commentService *CommentService) *PresentService {
	return &PresentService{presentRepository: *presentRepository, voteService: *voteService, commentService: *commentService}
}

func (service *PresentService) GetPresentListViewModel(eventId string) (*viewmodels.PresentListViewModel, error) {
	presents, err := service.presentRepository.GetPresentsByEventId(uuid.FromStringOrNil(eventId))
	if err != nil {
		logger.Log.Error(err, "Could not get presents")
		return nil, err
	}

	return service.mapPresentsToPresentListViewModel(eventId, presents), nil
}

func (service *PresentService) mapPresentsToPresentListViewModel(eventId string, presents []models.Present) *viewmodels.PresentListViewModel {
	var presentListItems []viewmodels.PresentListItemViewModel
	presentIds := service.extractIds(presents)

	presentIdToUpvoteCount, err := service.voteService.GetVoteCountMapByPresentIdsAndVoteType(presentIds, models.UPVOTE)
	if err != nil {
		logger.Log.Debug("Could not get upvote counts")
	}
	presentIdToDownvoteCount, err := service.voteService.GetVoteCountMapByPresentIdsAndVoteType(presentIds, models.DOWNVOTE)
	if err != nil {
		logger.Log.Debug("Could not get downvote counts")
	}

	for _, present := range presents {
		presentListItem := service.mapPresentToPresentListItemViewModel(&present, presentIdToUpvoteCount[present.Id.String()], presentIdToDownvoteCount[present.Id.String()])
		presentListItems = append(presentListItems, *presentListItem)
	}

	return &viewmodels.PresentListViewModel{
		EventId:  eventId,
		Presents: presentListItems,
	}
}

func (service *PresentService) GetPresentListItemViewModel(presentId string) (*viewmodels.PresentListItemViewModel, error) {
	present, err := service.presentRepository.GetPresentById(uuid.FromStringOrNil(presentId))
	if err != nil {
		logger.Log.Error(err, "Could not get present")
		return nil, err
	}

	return service.mapPresentToPresentListItemViewModel(present, 0, 0), nil
}

func (service *PresentService) mapPresentToPresentListItemViewModel(present *models.Present, upvoteCount int32, downvoteCount int32) *viewmodels.PresentListItemViewModel {
	dateLayout := "January 02, 2006"

	return &viewmodels.PresentListItemViewModel{
		Id:            present.Id.String(),
		Name:          present.Name,
		Link:          present.Link,
		UpvoteCount:   upvoteCount,
		DownvoteCount: downvoteCount,
		CreatedBy:     present.CreatedBy,
		CreatedAt:     present.CreatedAt.Format(dateLayout),
		Comments: viewmodels.CommentListViewModel{
			PresentId: present.Id.String(),
		},
	}
}

func (service *PresentService) CreatePresent(eventId string, createPresentRequest request.CreatePresentRequest) (uuid.UUID, error) {
	present := service.mapCreatePresentRequestToPresent(eventId, createPresentRequest)
	presentId, err := service.presentRepository.CreatePresent(present)
	if err != nil {
		logger.Log.Error(err, "Could not create present")
		return uuid.Nil, err
	}

	return presentId, err
}

func (service *PresentService) mapCreatePresentRequestToPresent(eventId string, createPresentRequest request.CreatePresentRequest) models.Present {
	createdUpdated := models.CreatedUpdated{
		CreatedBy: createPresentRequest.Username,
		UpdatedBy: createPresentRequest.Username,
	}

	return models.Present{
		EventId:        uuid.FromStringOrNil(eventId),
		Name:           createPresentRequest.Name,
		Link:           createPresentRequest.Link,
		CreatedUpdated: createdUpdated,
	}
}

func (service *PresentService) extractIds(presents []models.Present) []string {
	var ids []string
	for _, present := range presents {
		ids = append(ids, present.Id.String())
	}
	return ids
}
