package response

type GetEventResponse struct {
	Name      string   `json:"name"`
	Owner     string   `json:"owner"`
	Recipient string   `json:"recipient"`
	Users     []string `json:"users"`
}

type CreateEventResponse struct {
	Name string `json:"name"`
}
