package models

import uuid "github.com/satori/go.uuid"

type Event struct {
	Id        uuid.UUID
	Name      string
	Recipient string
	CreatedUpdated
}
