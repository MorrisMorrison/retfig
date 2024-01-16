package request

type CreateEventRequest struct {
	Name      string `json:"event-name"`
	Username  string `json:"username"`
	Recipient string `json:"recipient"`
}
