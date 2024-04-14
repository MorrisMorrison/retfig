package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Present struct {
	Id        uuid.UUID
	EventId   uuid.UUID
	Name      string
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
