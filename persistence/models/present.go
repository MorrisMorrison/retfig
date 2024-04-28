package models

import (
	uuid "github.com/satori/go.uuid"
)

type Present struct {
	Id      uuid.UUID
	EventId uuid.UUID
	Name    string
	Link    string
	CreatedUpdated
}

func NewPresent(eventId uuid.UUID, name string, link string, user string) *Present {
	createdUpdated := NewCreatedUpdated(user)

	return &Present{
		EventId:        eventId,
		Name:           name,
		Link:           link,
		CreatedUpdated: *createdUpdated,
	}
}
