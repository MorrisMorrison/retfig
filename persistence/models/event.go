package models

import uuid "github.com/satori/go.uuid"

type Event struct {
	Id        uuid.UUID
	Name      string
	Recipient string
	CreatedUpdated
}

func NewEvent(name string, recipient string, user string) *Event {
	createdUpdated := NewCreatedUpdated(user)

	return &Event{
		Name:           name,
		Recipient:      recipient,
		CreatedUpdated: *createdUpdated,
	}
}
