package services

import (
	"fmt"

	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/logger"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
	"github.com/MorrisMorrison/retfig/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type EventService struct {
	eventRepository repositories.EventRepository
	presentService  PresentService
}

func NewEventService(eventRepository *repositories.EventRepository, presentService *PresentService) *EventService {
	return &EventService{eventRepository: *eventRepository, presentService: *presentService}
}

func (service *EventService) CreateEvent(createEventRequest request.CreateEventRequest) (uuid.UUID, error) {
	event := service.mapCreateEventRequestToEvent(createEventRequest)
	eventId, err := service.eventRepository.CreateEvent(event)
	if err != nil {
		logger.Log.Error(err, "Could not create event")
		return uuid.Nil, err
	}

	return eventId, err
}

func (service *EventService) GetEventViewModel(id string) (*viewmodels.GetEventViewModel, error) {
	event, err := service.eventRepository.GetEventById(uuid.FromStringOrNil(id))
	if err != nil {
		logger.Log.Error(err, "Could not get event")
		return nil, err
	}

	participants, err := service.eventRepository.GetParticipantsByEventId(uuid.FromStringOrNil(id))
	if err != nil {
		logger.Log.Error(err, "Could not get participants")
		return nil, err
	}

	presents, err := service.presentService.GetPresentListViewModel(id)
	if err != nil {
		logger.Log.Error(err, "Could not get presents")
		return nil, err
	}

	return service.mapToGetEventViewModel(event, participants, *presents), nil
}

func (service *EventService) CreateParticipant(eventId string, createParticipantRequest request.CreateParticipantRequest) error {
	participant, err := service.eventRepository.GetParticipantByNameAndEventId(createParticipantRequest.Username, uuid.FromStringOrNil(eventId))
	if err != nil {
		return err
	}

	if participant != nil {
		return nil
	}

	event, err := service.eventRepository.GetEventById(uuid.FromStringOrNil(eventId))
	if err != nil {
		return err
	}

	if event.CreatedBy == createParticipantRequest.Username {
		return nil
	}

	fmt.Println("not found participant")

	return service.eventRepository.CreateParticipant(eventId, createParticipantRequest.Username)
}

func (service *EventService) mapCreateEventRequestToEvent(createEventRequest request.CreateEventRequest) models.Event {
	createdUpdated := models.CreatedUpdated{
		CreatedBy: createEventRequest.Username,
		UpdatedBy: createEventRequest.Username,
	}

	event := models.Event{
		Name:           createEventRequest.Name,
		Recipient:      createEventRequest.Recipient,
		CreatedUpdated: createdUpdated,
	}

	return event
}

func (service *EventService) mapToGetEventViewModel(event *models.Event, participants []models.Participant, presents viewmodels.PresentListViewModel) *viewmodels.GetEventViewModel {
	viewModel := &viewmodels.GetEventViewModel{
		Name:           event.Name,
		Recipient:      event.Recipient,
		CreatedBy:      event.CreatedBy,
		InvitationLink: service.getInvitationLink(event.Id),
		Participants:   service.extractUsernames(participants),
		Presents:       presents,
	}

	return viewModel
}

func (service *EventService) extractUsernames(participants []models.Participant) []string {
	var usernames []string
	for _, participant := range participants {
		usernames = append(usernames, participant.Name)
	}

	return usernames
}

func (service *EventService) getInvitationLink(eventId uuid.UUID) string {
	return "http://localhost:8080/events/" + eventId.String() + "/invitation"

}
