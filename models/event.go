package models

import uuid "github.com/satori/go.uuid"

type Event struct {
	Id        uuid.UUID
	Name      string
	Creator   string
	Recipient string
	Timestamps
}
