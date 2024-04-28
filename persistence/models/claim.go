package models

import (
	uuid "github.com/satori/go.uuid"
)

type Claim struct {
	PresentId uuid.UUID
	CreatedUpdated
}

func NewClaim(presentId string, user string) *Claim {
	createdUpdated := NewCreatedUpdated(user)

	return &Claim{
		PresentId:      uuid.FromStringOrNil(presentId),
		CreatedUpdated: *createdUpdated,
	}
}
