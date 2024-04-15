package request

type CreatePresentRequest struct {
	EventId  string `json:"eventId"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Username string `json:"username"`
}
