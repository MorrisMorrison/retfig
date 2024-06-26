package services

import (
	"sort"

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
	claimService      ClaimService
}

func NewPresentService(presentRepository *repositories.PresentRepository, voteService *VoteService, commentService *CommentService, claimService *ClaimService) *PresentService {
	return &PresentService{presentRepository: *presentRepository, voteService: *voteService, commentService: *commentService, claimService: *claimService}
}

func (service *PresentService) GetPresentListViewModel(eventId string, user string) (*viewmodels.PresentListViewModel, error) {
	presents, err := service.presentRepository.GetPresentsByEventId(uuid.FromStringOrNil(eventId))
	if err != nil {
		logger.LOG.Error(err, "Could not get presents")
		return nil, err
	}

	return service.mapPresentsToPresentListViewModel(eventId, presents, user), nil
}

func (service *PresentService) mapPresentsToPresentListViewModel(eventId string, presents []*models.Present, user string) *viewmodels.PresentListViewModel {
	var presentListItems []*viewmodels.PresentListItemViewModel
	presentIds := service.extractIds(presents)

	presentIdToUpvoteCount, err := service.voteService.GetVoteCountByPresentIdsAndVoteType(presentIds, models.UPVOTE)
	if err != nil {
		logger.LOG.Debug("Could not get upvote counts")
	}

	presentIdToDownvoteCount, err := service.voteService.GetVoteCountByPresentIdsAndVoteType(presentIds, models.DOWNVOTE)
	if err != nil {
		logger.LOG.Debug("Could not get downvote counts")
	}

	presentIdToCommentCount, err := service.commentService.GetCommentCountByPresentIds(presentIds)
	if err != nil {
		logger.LOG.Debug("Could not get comment counts")
	}

	presentIdToClaim, err := service.claimService.GetClaimsByPresentIds(presentIds)
	if err != nil {
		logger.LOG.Debug("Could not get claims")
	}

	presentIdToUserVotes, err := service.voteService.GetVotesByPresentIdsAndUser(presentIds, user)
	if err != nil {
		logger.LOG.Debug("Could not get claims")
	}

	for _, present := range presents {
		userVote := presentIdToUserVotes[present.Id.String()]
		isUpvotedByUser := userVote != nil && userVote.Type == models.UPVOTE
		isDownvotedByUser := userVote != nil && userVote.Type == models.DOWNVOTE

		presentListItem := service.mapPresentToPresentListItemViewModel(
			present,
			presentIdToClaim[present.Id.String()],
			presentIdToUpvoteCount[present.Id.String()],
			presentIdToDownvoteCount[present.Id.String()],
			presentIdToCommentCount[present.Id.String()],
			isUpvotedByUser,
			isDownvotedByUser)

		presentListItems = append(presentListItems, presentListItem)
	}

	sort.Slice(presentListItems, func(i, j int) bool {
		return presentListItems[i].UpvoteCount > presentListItems[j].UpvoteCount
	})

	return viewmodels.NewPresentListViewModel(eventId, presentListItems)
}

func (service *PresentService) GetSimplePresentListItemViewModel(presentId string) (*viewmodels.PresentListItemViewModel, error) {
	present, err := service.presentRepository.GetPresentById(uuid.FromStringOrNil(presentId))
	if err != nil {
		logger.LOG.Error(err, "Could not get present")
		return nil, err
	}

	return service.mapPresentToPresentListItemViewModel(present, nil, 0, 0, 0, false, false), nil
}

func (service *PresentService) GetPresentListItemViewModel(presentId string, user string) (*viewmodels.PresentListItemViewModel, error) {
	present, err := service.presentRepository.GetPresentById(uuid.FromStringOrNil(presentId))
	if err != nil {
		logger.LOG.Error(err, "Could not get present")
		return nil, err
	}

	claim, err := service.claimService.GetClaimByPresentId(presentId)
	if err != nil {
		logger.LOG.Error(err, "Could not get claim")
		return nil, err
	}

	upvoteCount, err := service.voteService.GetVoteCountByPresentIdAndVoteType(presentId, models.UPVOTE)
	if err != nil {
		logger.LOG.Error(err, "Could not get upvoteCount")
	}

	downvoteCount, err := service.voteService.GetVoteCountByPresentIdAndVoteType(presentId, models.DOWNVOTE)
	if err != nil {
		logger.LOG.Error(err, "Could not get downvoteCount")
	}

	commentCount, err := service.commentService.GetCommentCountByPresentId(presentId)
	if err != nil {
		logger.LOG.Error(err, "Could not get downvoteCount")
	}

	userVote, err := service.voteService.GetVoteByPresentIdAndUser(presentId, user)
	if err != nil {
		logger.LOG.Error(err, "Could not get vote")
	}

	isUpvotedByUser := userVote != nil && userVote.Type == models.UPVOTE
	isDownvotedByUser := userVote != nil && userVote.Type == models.DOWNVOTE

	return service.mapPresentToPresentListItemViewModel(present, claim, upvoteCount, downvoteCount, commentCount, isUpvotedByUser, isDownvotedByUser), nil
}

func (service *PresentService) mapPresentToPresentListItemViewModel(
	present *models.Present,
	claim *models.Claim,
	upvoteCount int32,
	downvoteCount int32,
	commentCount int32,
	isUpvotedByUser bool,
	isDownvotedByUser bool) *viewmodels.PresentListItemViewModel {

	dateLayout := "January 02, 2006"
	comments := &viewmodels.CommentListViewModel{}
	claimedBy := ""
	if claim != nil {
		claimedBy = claim.CreatedBy
	}

	return viewmodels.NewPresentListItemViewModel(
		present.Id.String(),
		present.EventId.String(),
		present.Name,
		present.Link,
		upvoteCount,
		downvoteCount,
		commentCount,
		present.CreatedBy,
		present.CreatedAt.Format(dateLayout),
		comments,
		claim != nil,
		claimedBy,
		isUpvotedByUser,
		isDownvotedByUser,
	)
}

func (service *PresentService) CreatePresent(eventId string, user string, createPresentRequest request.CreatePresentRequest) (uuid.UUID, error) {
	present := service.mapCreatePresentRequestToPresent(eventId, user, createPresentRequest)
	presentId, err := service.presentRepository.CreatePresent(present)
	if err != nil {
		logger.LOG.Error(err, "Could not create present")
		return uuid.Nil, err
	}

	return presentId, err
}

func (service *PresentService) mapCreatePresentRequestToPresent(eventId string, user string, createPresentRequest request.CreatePresentRequest) models.Present {
	return *models.NewPresent(uuid.FromStringOrNil(eventId), createPresentRequest.Name, createPresentRequest.Link, user)
}

func (service *PresentService) extractIds(presents []*models.Present) []string {
	var ids []string
	for _, present := range presents {
		ids = append(ids, present.Id.String())
	}
	return ids
}

func (service *PresentService) GetPresentById(presentId string) (*models.Present, error) {
	return service.presentRepository.GetPresentById(uuid.FromStringOrNil(presentId))
}
