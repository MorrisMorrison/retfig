package models

import uuid "github.com/satori/go.uuid"

type Participant struct {
	EventId uuid.UUID
	Name    string
	CreatedUpdated
}

func NewParticipant(eventId uuid.UUID, name string, user string) *Participant {
	createdUpdated := NewCreatedUpdated(user)

	return &Participant{
		EventId:        eventId,
		Name:           name,
		CreatedUpdated: *createdUpdated,
	}
}
