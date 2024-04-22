package request

type CreateCommentRequest struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}
