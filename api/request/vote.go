package request

type CreateVoteRequest struct {
	EventId   string `json:"eventId"`
	PresentId string `json:"presentId"`
	VoteType  string `json:"voteType"`
	Username  string `json:"username"`
}
