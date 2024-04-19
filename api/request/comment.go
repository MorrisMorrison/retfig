package request

type CreateCommentRequest struct {
	EventId   string `json:"eventId"`
	PresentId string `json:"presentId"`
	Content   string `json:"content"`
	Username  string `json:"username"`
}
