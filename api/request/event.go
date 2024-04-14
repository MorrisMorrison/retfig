package request

type CreateEventRequest struct {
	Name      string `json:"event-name"`
	Username  string `json:"username"`
	Recipient string `json:"recipient"`
}

type CreateParticipantRequest struct {
	Username string `json:"username"`
}

type UpdateEventRequest struct {
	Name      string `json:"event-name"`
	Recipient string `json:"recipient"`
}
