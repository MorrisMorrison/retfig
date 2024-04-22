package request

type CreatePresentRequest struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Username string `json:"username"`
}

type ClaimPresentRequest struct {
	Username string `json:"name"`
}
