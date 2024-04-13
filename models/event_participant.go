package models

import uuid "github.com/satori/go.uuid"

type EventParticipant struct {
	Id          uuid.UUID
	Event       Event
	Participant User
	Timestamps
}
