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
	event := mapCreateEventRequestToEvent(createEventRequest)
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

	return mapToGetEventViewModel(event, participants, *presents), nil
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

	if event.Creator == createParticipantRequest.Username {
		return nil
	}

	fmt.Println("not found participant")

	return service.eventRepository.CreateParticipant(eventId, createParticipantRequest.Username)
}

func mapCreateEventRequestToEvent(createEventRequest request.CreateEventRequest) models.Event {
	event := models.Event{
		Name:      createEventRequest.Name,
		Creator:   createEventRequest.Username,
		Recipient: createEventRequest.Recipient,
	}

	return event
}

func mapToGetEventViewModel(event *models.Event, participants []models.EventParticipant, presents viewmodels.PresentListViewModel) *viewmodels.GetEventViewModel {
	viewModel := &viewmodels.GetEventViewModel{
		Name:           event.Name,
		Recipient:      event.Recipient,
		Creator:        event.Creator,
		InvitationLink: getInvitationLink(event.Id),
		Participants:   extractUsernames(participants),
		Presents:       presents,
	}

	return viewModel
}

func extractUsernames(participants []models.EventParticipant) []string {
	var usernames []string
	for _, participant := range participants {
		usernames = append(usernames, participant.Participant)
	}

	return usernames
}

func getInvitationLink(eventId uuid.UUID) string {
	return "http://localhost:8080/events/" + eventId.String() + "/invitation"

}
