package models

import uuid "github.com/satori/go.uuid"

type Comment struct {
	PresentId uuid.UUID
	Content   string
	CreatedUpdated
}
