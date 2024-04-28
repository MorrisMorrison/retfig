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

func NewVote(presentId uuid.UUID, voteType VoteType, user string) *Vote {
	createdUpdated := NewCreatedUpdated(user)

	return &Vote{
		PresentId:      presentId,
		Type:           voteType,
		CreatedUpdated: *createdUpdated,
	}
}
