package models

import uuid "github.com/satori/go.uuid"

type Comment struct {
	PresentId uuid.UUID
	Content   string
	CreatedUpdated
}

func NewComment(presentId uuid.UUID, content string, user string) *Comment {
	createdUpdated := NewCreatedUpdated(user)

	return &Comment{
		PresentId:      presentId,
		Content:        content,
		CreatedUpdated: *createdUpdated,
	}
}
