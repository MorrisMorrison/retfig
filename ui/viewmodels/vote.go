package viewmodels

import "github.com/MorrisMorrison/retfig/persistence/models"

type VoteButtonViewModel struct {
	EventId       string
	PresentId     string
	VoteType      models.VoteType
	VoteCount     int32
	CreatedBy     string
	IsVotedByUser bool
}

func NewVoteButtonViewModel(eventId string, presentId string, voteType models.VoteType, voteCount int32, createdBy string, isVotedByUser bool) *VoteButtonViewModel {
	return &VoteButtonViewModel{
		EventId:       eventId,
		PresentId:     presentId,
		VoteType:      voteType,
		VoteCount:     voteCount,
		CreatedBy:     createdBy,
		IsVotedByUser: isVotedByUser,
	}
}

type VoteButtonsViewModel struct {
	EventId           string
	PresentId         string
	UpvoteCount       int32
	DownvoteCount     int32
	CreatedBy         string
	IsUpvotedByUser   bool
	IsDownvotedByUser bool
}

func NewVoteButtonsViewModel(eventId string, presentId string, upvoteCount int32, downvoteCount int32, createdBy string, isUpvotedByUser bool, isDownvotedByUser bool) *VoteButtonsViewModel {
	return &VoteButtonsViewModel{
		EventId:           eventId,
		PresentId:         presentId,
		UpvoteCount:       upvoteCount,
		DownvoteCount:     downvoteCount,
		CreatedBy:         createdBy,
		IsUpvotedByUser:   isUpvotedByUser,
		IsDownvotedByUser: isDownvotedByUser,
	}
}
