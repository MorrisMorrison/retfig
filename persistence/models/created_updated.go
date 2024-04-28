package models

import "time"

type CreatedUpdated struct {
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCreatedUpdated(user string) *CreatedUpdated {
	return &CreatedUpdated{
		CreatedBy: user,
		UpdatedBy: user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
