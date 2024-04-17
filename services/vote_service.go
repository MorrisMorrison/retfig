package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
	uuid "github.com/satori/go.uuid"
)

type VoteService struct {
	voteRepository repositories.VoteRepository
}

func NewVoteService(voteRepository *repositories.VoteRepository) *VoteService {
	return &VoteService{voteRepository: *voteRepository}
}

func (voteService *VoteService) CreateVote(request request.CreateVoteRequest) error {
	vote := voteService.mapCreateVoteRequestToVote(request)
	return voteService.voteRepository.CreateVote(vote)
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
