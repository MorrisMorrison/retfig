package request

type CreateVoteRequest struct {
	VoteType string `json:"voteType"`
	Username string `json:"username"`
}
