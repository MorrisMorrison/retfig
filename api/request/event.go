package request

type CreateEventRequest struct {
	Name      string `json:"eventName"`
	Username  string `json:"username"`
	Recipient string `json:"recipient"`
}

type UpdateEventRequest struct {
	Name      string `json:"eventName"`
	Recipient string `json:"recipient"`
}
