package models

import uuid "github.com/satori/go.uuid"

type EventParticipant struct {
	EventId     uuid.UUID
	Participant string
	CreatedUpdated
}
