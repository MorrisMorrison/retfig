package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/infrastructure/logger"
	"github.com/MorrisMorrison/retfig/persistence/models"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	uuid "github.com/satori/go.uuid"
)

type ClaimService struct {
	claimRepository repositories.ClaimRepository
}

func NewClaimService(claimRepository *repositories.ClaimRepository) *ClaimService {
	return &ClaimService{claimRepository: *claimRepository}
}

func (claimService *ClaimService) CreateClaim(presentId string, request request.CreateClaimRequest) error {
	claim := claimService.mapCreateClaimRequestToClaim(presentId, request)
	return claimService.claimRepository.CreateClaim(claim)
}

func (claimService *ClaimService) mapCreateClaimRequestToClaim(presentId string, request request.CreateClaimRequest) models.Claim {
	createdUpdated := models.CreatedUpdated{
		CreatedBy: request.Username,
		UpdatedBy: request.Username,
	}

	return models.Claim{
		PresentId:      uuid.FromStringOrNil(presentId),
		CreatedUpdated: createdUpdated,
	}
}

func (claimService *ClaimService) DeleteClaim(presentId string) error {
	err := claimService.claimRepository.DeleteClaimByPresentId(uuid.FromStringOrNil(presentId))
	if err != nil {
		logger.LOG.Error(err, "Could not delete existing vote")
		return err
	}

	return nil
}
