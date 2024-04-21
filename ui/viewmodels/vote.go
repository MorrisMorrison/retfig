package viewmodels

import "github.com/MorrisMorrison/retfig/persistence/models"

type VoteButtonViewModel struct {
	EventId   string
	PresentId string
	VoteType  models.VoteType
	VoteCount int32
	CreatedBy string
}

func NewVoteButtonViewModel(eventId string, presentId string, voteType models.VoteType, voteCount int32, createdBy string) *VoteButtonViewModel {
	return &VoteButtonViewModel{
		EventId:   eventId,
		PresentId: presentId,
		VoteType:  voteType,
		VoteCount: voteCount,
		CreatedBy: createdBy,
	}
}

type VoteButtonsViewModel struct {
	EventId       string
	PresentId     string
	UpvoteCount   int32
	DownvoteCount int32
	CreatedBy     string
}

func NewVoteButtonsViewModel(eventId string, presentId string, upvoteCount int32, downvoteCount int32, createdBy string) *VoteButtonsViewModel {
	return &VoteButtonsViewModel{
		EventId:       eventId,
		PresentId:     presentId,
		UpvoteCount:   upvoteCount,
		DownvoteCount: downvoteCount,
		CreatedBy:     createdBy,
	}
}
