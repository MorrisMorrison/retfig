package models

import uuid "github.com/satori/go.uuid"

type VoteType string

const (
	UPVOTE   VoteType = "UPVOTE"
	DOWNVOTE VoteType = "DOWNVOTE"
)

type Vote struct {
	PresentId uuid.UUID
	Type      VoteType
	CreatedUpdated
}
