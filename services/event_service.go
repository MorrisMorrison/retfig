package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/logger"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
	"github.com/MorrisMorrison/retfig/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type EventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) *EventService {
	return &EventService{eventRepository: *eventRepository}
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

func (service *EventService) GetEventViewModel(id uuid.UUID) (*viewmodels.GetEventViewModel, error) {
	event, err := service.eventRepository.GetEventById(id)
	if err != nil {
		logger.Log.Error(err, "Could not get event")
		return nil, err
	}

	return mapEventToGetEventViewModel(event), nil
}

func mapCreateEventRequestToEvent(createEventRequest request.CreateEventRequest) models.Event {
	event := models.Event{
		Name:         createEventRequest.Name,
		CreatorEmail: createEventRequest.Username,
		Recipient:    createEventRequest.Recipient,
	}

	return event
}

func mapEventToGetEventViewModel(event *models.Event) *viewmodels.GetEventViewModel {
	viewModel := &viewmodels.GetEventViewModel{
		Name:         event.Name,
		Recipient:    event.Recipient,
		CreatorEmail: event.CreatorEmail,
	}

	return viewModel
}
