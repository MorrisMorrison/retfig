package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/infrastructure/logger"
	"github.com/MorrisMorrison/retfig/persistence/models"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	"github.com/MorrisMorrison/retfig/ui/viewmodels"
	"github.com/MorrisMorrison/retfig/utils/links"
	uuid "github.com/satori/go.uuid"
)

type EventService struct {
	eventRepository    repositories.EventRepository
	presentService     PresentService
	participantService ParticipantService
}

func NewEventService(eventRepository *repositories.EventRepository, presentService *PresentService, participantService *ParticipantService) *EventService {
	return &EventService{eventRepository: *eventRepository, presentService: *presentService, participantService: *participantService}
}

func (service *EventService) CreateEvent(createEventRequest request.CreateEventRequest) (uuid.UUID, error) {
	event := service.mapCreateEventRequestToEvent(createEventRequest)
	eventId, err := service.eventRepository.CreateEvent(event)
	if err != nil {
		logger.LOG.Error(err, "Could not create event")
		return uuid.Nil, err
	}

	return eventId, err
}

func (service *EventService) GetEventViewModel(eventId string, user string) (*viewmodels.GetEventViewModel, error) {
	event, err := service.eventRepository.GetEventById(uuid.FromStringOrNil(eventId))
	if err != nil {
		logger.LOG.Error(err, "Could not get event")
		return nil, err
	}

	participants, err := service.participantService.GetParticipantsByEventId(eventId)
	if err != nil {
		logger.LOG.Error(err, "Could not get participants")
		return nil, err
	}

	presents, err := service.presentService.GetPresentListViewModel(eventId, user)
	if err != nil {
		logger.LOG.Error(err, "Could not get presents")
		return nil, err
	}

	return service.mapToGetEventViewModel(event, participants, presents), nil
}

func (service *EventService) mapCreateEventRequestToEvent(createEventRequest request.CreateEventRequest) models.Event {
	return *models.NewEvent(createEventRequest.Name, createEventRequest.Recipient, createEventRequest.Username)
}

func (service *EventService) mapToGetEventViewModel(event *models.Event, participants []*models.Participant, presents *viewmodels.PresentListViewModel) *viewmodels.GetEventViewModel {
	return viewmodels.NewGetEventViewModel(event.Name, event.CreatedBy, event.Recipient, links.BuildInvitationLink(event.Id.String()), service.extractUsernames(participants), presents)
}

func (service *EventService) extractUsernames(participants []*models.Participant) []string {
	var usernames []string
	for _, participant := range participants {
		usernames = append(usernames, participant.Name)
	}

	return usernames
}

func (service *EventService) GetEventById(eventId string) (*models.Event, error) {
	return service.eventRepository.GetEventById(uuid.FromStringOrNil(eventId))
}
