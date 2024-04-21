package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/infrastructure/logger"
	"github.com/MorrisMorrison/retfig/persistence/models"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	"github.com/MorrisMorrison/retfig/ui/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type VoteService struct {
	voteRepository repositories.VoteRepository
}

func NewVoteService(voteRepository *repositories.VoteRepository) *VoteService {
	return &VoteService{voteRepository: *voteRepository}
}

func (voteService *VoteService) GetVoteButtonsViewModel(eventId string, createVoteRequest request.CreateVoteRequest) (*viewmodels.VoteButtonsViewModel, error) {
	presentId := uuid.FromStringOrNil(createVoteRequest.PresentId)

	vote, err := voteService.voteRepository.GetVoteByPresentIdAndUser(presentId, createVoteRequest.Username)
	if err != nil {
		return nil, err
	}

	upvoteCount, err := voteService.voteRepository.GetVoteCountByPresentIdAndVoteType(presentId, models.UPVOTE)
	if err != nil {
		logger.Log.Debug("Could not fetch vote count")
	}

	downvoteCount, err := voteService.voteRepository.GetVoteCountByPresentIdAndVoteType(presentId, models.DOWNVOTE)
	if err != nil {
		logger.Log.Debug("Could not fetch vote count")
	}

	voteButtonsViewModel := viewmodels.NewVoteButtonsViewModel(eventId, vote.PresentId.String(), upvoteCount, downvoteCount, vote.CreatedBy, vote.Type == models.UPVOTE, vote.Type == models.DOWNVOTE)
	return voteButtonsViewModel, nil
}

func (voteService *VoteService) GetVoteButtonViewModel(eventId string, createVoteRequest request.CreateVoteRequest) (*viewmodels.VoteButtonViewModel, error) {
	presentId := uuid.FromStringOrNil(createVoteRequest.PresentId)

	vote, err := voteService.voteRepository.GetVoteByPresentIdAndUser(presentId, createVoteRequest.Username)
	if err != nil {
		return nil, err
	}

	voteCount, err := voteService.voteRepository.GetVoteCountByPresentIdAndVoteType(presentId, models.VoteType(createVoteRequest.VoteType))
	if err != nil {
		logger.Log.Debug("Could not fetch vote count")
	}

	createVoteButtonViewModel := viewmodels.NewVoteButtonViewModel(eventId, vote.PresentId.String(), vote.Type, voteCount, vote.CreatedBy, true)
	return createVoteButtonViewModel, nil
}

func (voteService *VoteService) CreateVote(request request.CreateVoteRequest) error {
	err := voteService.deleteVoteIfExists(request.PresentId, request.Username, request.VoteType)
	if err != nil {
		return err
	}

	vote := voteService.mapCreateVoteRequestToVote(request)
	return voteService.voteRepository.CreateVote(vote)
}

func (voteService *VoteService) deleteVoteIfExists(presentId string, username string, voteType string) error {
	existingVote, err := voteService.voteRepository.GetVoteByPresentIdAndUser(uuid.FromStringOrNil(presentId), username)
	if err != nil {
		logger.Log.Error(err, "Could not get existing vote")
		return err
	}

	shouldDeleteExistingVote := existingVote != nil && existingVote.Type != models.VoteType(voteType)
	if shouldDeleteExistingVote {
		err := voteService.voteRepository.DeleteVoteByPresentIdAndUsername(uuid.FromStringOrNil(presentId), username)
		if err != nil {
			logger.Log.Error(err, "Could not delete existing vote")
			return err
		}
	}

	return nil
}

func (voteService *VoteService) GetVoteCountMapByPresentIdsAndVoteType(presentIds []string, voteType models.VoteType) (map[string]int32, error) {
	return voteService.voteRepository.GetVoteCountMapByPresentIdsAndVoteType(presentIds, voteType)
}

func (voteService *VoteService) mapCreateVoteRequestToVote(request request.CreateVoteRequest) models.Vote {
	createdUpdated := models.CreatedUpdated{
		CreatedBy: request.Username,
		UpdatedBy: request.Username,
	}

	return models.Vote{
		PresentId:      uuid.FromStringOrNil(request.PresentId),
		Type:           models.VoteType(request.VoteType),
		CreatedUpdated: createdUpdated,
	}
}
