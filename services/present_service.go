package services

import (
	"github.com/MorrisMorrison/retfig/api/request"
	"github.com/MorrisMorrison/retfig/logger"
	"github.com/MorrisMorrison/retfig/models"
	"github.com/MorrisMorrison/retfig/repositories"
	"github.com/MorrisMorrison/retfig/viewmodels"
	uuid "github.com/satori/go.uuid"
)

type PresentService struct {
	presentRepository repositories.PresentRepository
}

func NewPresentService(presentRepository *repositories.PresentRepository) *PresentService {
	return &PresentService{presentRepository: *presentRepository}
}

func (service *PresentService) GetPresentListViewModel(eventId string) (*viewmodels.PresentListViewModel, error) {
	presents, err := service.presentRepository.GetPresentsByEventId(uuid.FromStringOrNil(eventId))
	if err != nil {
		logger.Log.Error(err, "Could not get presents")
		return nil, err
	}

	return mapPresentsToPresentListViewModel(eventId, presents), nil
}

func mapPresentsToPresentListViewModel(eventId string, presents []models.Present) *viewmodels.PresentListViewModel {
	var presentListItems []viewmodels.PresentListItemViewModel

	for _, present := range presents {
		presentListItem := mapPresentToPresentListItemViewModel(present)
		presentListItems = append(presentListItems, *presentListItem)
	}

	return &viewmodels.PresentListViewModel{
		EventId:  eventId,
		Presents: presentListItems,
	}
}

func mapPresentToPresentListItemViewModel(present models.Present) *viewmodels.PresentListItemViewModel {
	dateLayout := "January 02, 2006"

	return &viewmodels.PresentListItemViewModel{
		Name:      present.Name,
		Link:      present.Link,
		Creator:   present.Creator,
		CreatedAt: present.CreatedAt.Format(dateLayout),
	}
}

func (service *PresentService) CreatePresent(eventId string, createPresentRequest request.CreatePresentRequest) (uuid.UUID, error) {
	present := mapCreatePresentRequestToPresent(eventId, createPresentRequest)
	presentId, err := service.presentRepository.CreatePresent(present)
	if err != nil {
		logger.Log.Error(err, "Could not create present")
		return uuid.Nil, err
	}

	return presentId, err
}

func mapCreatePresentRequestToPresent(eventId string, createPresentRequest request.CreatePresentRequest) models.Present {
	return models.Present{
		EventId: uuid.FromStringOrNil(eventId),
		Name:    createPresentRequest.Name,
		Link:    createPresentRequest.Link,
		Creator: createPresentRequest.Username,
	}
}
