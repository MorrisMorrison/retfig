package models

import uuid "github.com/satori/go.uuid"

type Claim struct {
	PresentId uuid.UUID
	CreatedUpdated
}
