package models

import uuid "github.com/satori/go.uuid"

type Participant struct {
	EventId uuid.UUID
	Name    string
	CreatedUpdated
}
