package services

type ResourceAcessService struct {
	eventService       *EventService
	presentService     *PresentService
	participantService *ParticipantService
}

func NewResourceAccessService(eventService *EventService, presentService *PresentService, participantService *ParticipantService) *ResourceAcessService {
	return &ResourceAcessService{
		eventService:       eventService,
		presentService:     presentService,
		participantService: participantService,
	}
}

func (service *ResourceAcessService) CanAccessEvent(eventId string, user string) (bool, error) {
	event, err := service.eventService.GetEventById(eventId)
	if err != nil {
		return false, err
	}

	if event == nil {
		return false, nil
	}

	if event.CreatedBy == user {
		return true, nil
	}

	participant, err := service.participantService.GetParticipantByNameAndEventId(eventId, user)
	if err != nil {
		return false, err
	}

	if participant == nil {
		return false, nil
	}

	return true, nil
}

func (service *ResourceAcessService) CanAccessPresent(presentId string, user string) (bool, error) {
	present, err := service.presentService.GetPresentById(presentId)
	if err != nil {
		return false, err
	}

	if present == nil {
		return false, nil
	}

	if present.CreatedBy == user {
		return true, nil
	}

	return service.CanAccessEvent(present.EventId.String(), user)
}
