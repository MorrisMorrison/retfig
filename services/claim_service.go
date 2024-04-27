package services

import (
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

func (claimService *ClaimService) CreateClaim(presentId string, user string) error {
	claim := models.NewClaim(presentId, user)
	return claimService.claimRepository.CreateClaim(*claim)
}

func (claimService *ClaimService) DeleteClaim(presentId string) error {
	err := claimService.claimRepository.DeleteClaimByPresentId(uuid.FromStringOrNil(presentId))
	if err != nil {
		logger.LOG.Error(err, "Could not delete existing vote")
		return err
	}

	return nil
}

func (claimService *ClaimService) GetClaimsByPresentIds(presentIds []string) (map[string]*models.Claim, error) {
	return claimService.claimRepository.GetClaimsByPresentIds(presentIds)
}

func (claimService *ClaimService) GetClaimByPresentId(presentId string) (*models.Claim, error) {
	return claimService.claimRepository.GetClaimByPresentId(uuid.FromStringOrNil(presentId))
}
