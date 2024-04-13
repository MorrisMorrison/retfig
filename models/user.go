package models

import uuid "github.com/satori/go.uuid"

type User struct {
	Id               uuid.UUID
	Name             string
	Email            string
	IsActivated      bool
	IsEmailConfirmed bool
	Timestamps
}
