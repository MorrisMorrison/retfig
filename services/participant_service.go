package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
	uuid "github.com/satori/go.uuid"
)

type ParticipantService struct {
	participantRepository repositories.ParticipantRepository
}

func NewParticipantService(participantRepository *repositories.ParticipantRepository) *ParticipantService {
	return &ParticipantService{participantRepository: *participantRepository}
}

func (service *ParticipantService) GetParticipantsByEventId(eventId string) ([]models.Participant, error) {
	return service.participantRepository.GetParticipantsByEventId(uuid.FromStringOrNil(eventId))
}

func (service *ParticipantService) CreateParticipant(eventId string, createParticipantRequest request.CreateParticipantRequest) error {
	participant, err := service.participantRepository.GetParticipantByNameAndEventId(createParticipantRequest.Username, uuid.FromStringOrNil(eventId))
	if err != nil {
		return err
	}

	if participant != nil {
		return nil
	}

	// event, err := service.eventRepository.GetEventById(uuid.FromStringOrNil(eventId))
	// if err != nil {
	// 	return err
	// }

	// if event.CreatedBy == createParticipantRequest.Username {
	// 	return nil
	// }

	return service.participantRepository.CreateParticipant(eventId, createParticipantRequest.Username)
}
