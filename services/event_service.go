package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
)

type EventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) *EventService {
	return &EventService{eventRepository: *eventRepository}
}

func (service *EventService) CreateEvent(createEventRequest request.CreateEventRequest) (models.Event, error) {
	event := mapCreateEventRequestToEvent(createEventRequest)
	err := service.eventRepository.CreateEvent(event)
	return event, err
}

func mapCreateEventRequestToEvent(createEventRequest request.CreateEventRequest) models.Event {
	event := models.Event{
		Name: createEventRequest.Name,
	}

	return event
}
